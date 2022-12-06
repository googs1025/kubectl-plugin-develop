package main

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubectl_plugin_develop/common"
	"kubectl_plugin_develop/initClient"
	"log"
	"os"
	"strconv"
)

var ServiceCmd = &cobra.Command{}

func ServiceCommand() *cobra.Command {
	client := initClient.InitClient()

	ServiceCmd = &cobra.Command{
		Use:          "services [flags]",
		Short:        "list services",
		Example:      "kubectl services [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			err = ListServicesWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return ServiceCmd


}

func ListServicesWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	serviceList, err := client.CoreV1().Services(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Service名称", "服务发现类型", "Cluster IP", "NodePort", "TargetPort"}

	if common.ShowLabels {
		content = append(content, "标签")
	}

	table.SetHeader(content)


	for _, service := range serviceList.Items {
		// TODO: 还会有多组port的显示bug
		portStringList := []string{strconv.Itoa(int(service.Spec.Ports[0].NodePort)), strconv.Itoa(int(service.Spec.Ports[0].TargetPort.IntVal))}

		serviceRow := []string{service.Name, string(service.Spec.Type), service.Spec.ClusterIP}
		serviceRow = append(serviceRow, portStringList...)

		if common.ShowLabels {
			serviceRow = append(serviceRow, common.LabelsMapToString(service.Labels))
		}
		table.Append(serviceRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}


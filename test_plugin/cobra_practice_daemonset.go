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

var DaemonsetCmd = &cobra.Command{}

func DaemonsetCommand() *cobra.Command {
	client := initClient.InitClient()

	DaemonsetCmd = &cobra.Command{
		Use:          "daemonsets [flags]",
		Short:        "list daemonsets",
		Example:      "kubectl daemonsets [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			err = ListDaemonsetsWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return DaemonsetCmd


}

func ListDaemonsetsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	daemonsetList, err := client.AppsV1().DaemonSets(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Daemonset名称", "Namespace", "DesiredNumber", "CurrentNumber", "ReadyNumber", "AvailableNumber"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, daemonset := range daemonsetList.Items {
		daemonsetRow := []string{daemonset.Name, daemonset.Namespace, strconv.Itoa(int(daemonset.Status.DesiredNumberScheduled)), strconv.Itoa(int(daemonset.Status.CurrentNumberScheduled)), strconv.Itoa(int(daemonset.Status.NumberReady)),
			strconv.Itoa(int(daemonset.Status.NumberAvailable))}
		if common.ShowLabels {
			daemonsetRow = append(daemonsetRow, common.LabelsMapToString(daemonset.Labels))
		}
		if common.ShowAnnotations {
			daemonsetRow = append(daemonsetRow, common.AnnotationsMapToString(daemonset.Annotations))
		}


		table.Append(daemonsetRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}





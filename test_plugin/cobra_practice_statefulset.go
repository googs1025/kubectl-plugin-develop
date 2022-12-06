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

var StatefulsetCmd = &cobra.Command{}

func StatefulsetCommand() *cobra.Command {
	client := initClient.InitClient()

	StatefulsetCmd = &cobra.Command{
		Use:          "statefulsets [flags]",
		Short:        "list statefulsets",
		Example:      "kubectl statefulsets [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			err = ListStatefulsetsWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return StatefulsetCmd


}

func ListStatefulsetsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	statefulsetList, err := client.AppsV1().StatefulSets(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Statefulset名称", "Namespace", "副本数", "Available副本数", "Ready副本数"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, statefulset := range statefulsetList.Items {
		statefulsetRow := []string{statefulset.Name, statefulset.Namespace, strconv.Itoa(int(statefulset.Status.Replicas)), strconv.Itoa(int(statefulset.Status.AvailableReplicas)), strconv.Itoa(int(statefulset.Status.ReadyReplicas))}
		if common.ShowLabels {
			statefulsetRow = append(statefulsetRow, common.LabelsMapToString(statefulset.Labels))
		}
		if common.ShowAnnotations {
			statefulsetRow = append(statefulsetRow, common.AnnotationsMapToString(statefulset.Annotations))
		}


		table.Append(statefulsetRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}


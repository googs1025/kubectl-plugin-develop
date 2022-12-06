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

var ConfigmapCmd = &cobra.Command{}

func ConfigmapCommand() *cobra.Command {
	client := initClient.InitClient()

	ConfigmapCmd = &cobra.Command{
		Use:          "configmaps [flags]",
		Short:        "list configmaps",
		Example:      "kubectl configmaps [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}

			err = ListConfigmapsWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return ConfigmapCmd


}

func ListConfigmapsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	configmapList, err := client.CoreV1().ConfigMaps(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Configmap名称", "Namespace", "Data个数"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, configmap := range configmapList.Items {
		configmapRow := []string{configmap.Name, configmap.Namespace, strconv.Itoa(len(configmap.Data))}
		if common.ShowLabels {
			configmapRow = append(configmapRow, common.LabelsMapToString(configmap.Labels))
		}
		if common.ShowAnnotations {
			configmapRow = append(configmapRow, common.AnnotationsMapToString(configmap.Annotations))
		}


		table.Append(configmapRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}


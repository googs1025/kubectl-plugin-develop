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
)

var PodCmd = &cobra.Command{}

func PodCommand() *cobra.Command {
	client := initClient.InitClient()

	PodCmd = &cobra.Command{
		Use:          "pods [flags]",
		Short:        "list pods",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			err = ListPodsWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return PodCmd


}


func ListPodsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	podList, err := client.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"POD名称", "Namespace", "POD IP", "状态"}


	table.SetHeader(content)


	for _, pod := range podList.Items {
		podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}

		table.Append(podRow)
	}
	// 去掉表格线
	table = common.TableSet(table)

	table.Render()

	return nil


}


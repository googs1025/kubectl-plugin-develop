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

var ShowLabels bool

func init() {
	ShowLabels = false
}

func main() {

	client := initClient.InitClient()
	cmd := &cobra.Command{
		Use: "kubectl pods [flags]",
		Short: "list pods",
		Example: "kubectl pods [flags]",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ns, err := cmd.Flags().GetString("namespace")
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

	initClient.MergeFlags(cmd)

	//cmd.Flags().Bool("show-labels", false, "kubectl pods --show-labels")
	cmd.Flags().BoolVar(&ShowLabels, "show-labels", false, "kubectl pods --show-labels")

	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
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
	context := []string{"POD名称", "Namespace", "POD IP", "状态"}

	if ShowLabels {
		context = append(context, "标签")
	}
	table.SetHeader(context)

	//for _, pod := range podList.Items {
	//	fmt.Printf("pod Name:%s, pod Namespace:%s\n", pod.Name, pod.Namespace)
	//}

	for _, pod := range podList.Items {
		podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}

		if ShowLabels {
			podRow = append(podRow, common.MapToString(pod.Labels))
		}
		table.Append(podRow)
	}
	// 去掉表格线
	//table.SetAutoWrapText(false)
	//table.SetAutoFormatHeaders(true)
	//table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	//table.SetAlignment(tablewriter.ALIGN_LEFT)
	//table.SetCenterSeparator("")
	//table.SetColumnSeparator("")
	//table.SetRowSeparator("")
	//table.SetHeaderLine(false)
	//table.SetBorder(false)
	//table.SetTablePadding("\t") // pad with tabs
	//table.SetNoWhiteSpace(true)

	table.Render()

	return nil
}
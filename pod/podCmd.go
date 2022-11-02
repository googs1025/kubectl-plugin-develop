package pod

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubectl_plugin_develop/common"
	"kubectl_plugin_develop/initClient"
	"os"
	"context"
	"log"
)


func PodRunCmd(cmd *cobra.Command,args []string) error {

	client := initClient.InitClient()
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
}


func ListPodsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	podList, err := client.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"POD名称", "Namespace", "POD IP", "状态"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	table.SetHeader(content)

	//for _, pod := range podList.Items {
	//	fmt.Printf("pod Name:%s, pod Namespace:%s\n", pod.Name, pod.Namespace)
	//}

	for _, pod := range podList.Items {
		podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}

		if common.ShowLabels {
			podRow = append(podRow, common.MapToString(pod.Labels))
		}
		table.Append(podRow)
	}
	// 去掉表格线
	table = common.TableSet(table)

	table.Render()

	return nil
}

package pod

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
	"kubectl_plugin_develop/common"
	"os"
)

func ListAndWatchPodsWithNamespace(namespace string) error {
	podList,err:=fact.Core().V1().Pods().Lister().Pods(namespace).
		List(labels.Everything())
	if err != nil {
		return err
	}

	fmt.Println("从缓存取")
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"POD名称", "Namespace", "POD IP", "状态"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	table.SetHeader(content)

	for _, pod := range podList {
		podRow := []string{
			pod.Name,
			pod.Namespace,
			pod.Status.PodIP,
			string(pod.Status.Phase),
		}
		table.Append(podRow)
	}
	table = common.TableSet(table)
	table.Render()
	return nil
}


var cacheCmd = &cobra.Command{
	Use:          "cache",
	Short:        "pods by cache",
	Hidden:true,
	RunE: func(c *cobra.Command, args []string) error {

		ns, err := c.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		if ns == "" {
			ns = "default"
		}
		err = ListAndWatchPodsWithNamespace(ns)
		if err != nil {
			return err
		}

		return nil
	},

}


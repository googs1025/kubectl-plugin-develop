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

var JobCmd = &cobra.Command{}

func JobCommand() *cobra.Command {
	client := initClient.InitClient()

	JobCmd = &cobra.Command{
		Use:          "jobs [flags]",
		Short:        "list jobs",
		Example:      "kubectl jobs [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}

			err = ListJobsWithNamespace(client, ns)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return JobCmd


}

func ListJobsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	jobList, err := client.BatchV1().Jobs(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Job名称", "Completions", "Parallelism"}
	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, job := range jobList.Items {

		jobRow := []string{job.Name, strconv.Itoa(int(*job.Spec.Completions)), strconv.Itoa(int(*job.Spec.Parallelism))}
		if common.ShowLabels {
			jobRow = append(jobRow, common.LabelsMapToString(job.Labels))
		}
		if common.ShowAnnotations {
			jobRow = append(jobRow, common.AnnotationsMapToString(job.Annotations))
		}

		table.Append(jobRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}


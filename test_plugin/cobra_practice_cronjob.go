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
	"fmt"
)

var CronjobCmd = &cobra.Command{}


func CronjobCommand() *cobra.Command {
	client := initClient.InitClient()

	CronjobCmd = &cobra.Command{
		Use:          "cronjobs [flags]",
		Short:        "list cronjobs",
		Example:      "kubectl cronjobs [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				fmt.Println("111")
				return err
			}
			if ns == "" {
				ns = "default"
			}

			err = ListCronjobsWithNamespace(client, ns)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return CronjobCmd


}

func ListCronjobsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	cronjobList, err := client.BatchV1().CronJobs(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"CronJob名称", "Namespace","Completions", "Parallelism", "Schedule"}


	table.SetHeader(content)


	for _, cronjob := range cronjobList.Items {

		serviceRow := []string{cronjob.Name, cronjob.Namespace,strconv.Itoa(int(*cronjob.Spec.JobTemplate.Spec.Completions)), strconv.Itoa(int(*cronjob.Spec.JobTemplate.Spec.Parallelism)), cronjob.Spec.Schedule}


		table.Append(serviceRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}



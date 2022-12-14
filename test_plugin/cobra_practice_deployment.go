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

var DeploymentCmd = &cobra.Command{}

func DeploymentCommand() *cobra.Command {
	client := initClient.InitClient()

	DeploymentCmd = &cobra.Command{
		Use:          "deployments [flags]",
		Short:        "list deployments",
		Example:      "kubectl deployments [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			err = ListDeploymentsWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return DeploymentCmd


}

func ListDeploymentsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	deploymentList, err := client.AppsV1().Deployments(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Deployment名称", "Namespace", "副本数", "Available副本数", "Ready副本数"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, deployment := range deploymentList.Items {
		deploymentRow := []string{deployment.Name, deployment.Namespace, strconv.Itoa(int(deployment.Status.Replicas)), strconv.Itoa(int(deployment.Status.AvailableReplicas)), strconv.Itoa(int(deployment.Status.ReadyReplicas))}
		if common.ShowLabels {
			deploymentRow = append(deploymentRow, common.LabelsMapToString(deployment.Labels))
		}
		if common.ShowAnnotations {
			deploymentRow = append(deploymentRow, common.AnnotationsMapToString(deployment.Annotations))
		}


		table.Append(deploymentRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}


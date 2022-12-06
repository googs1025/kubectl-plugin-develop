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

var SecretCmd = &cobra.Command{}

func SecretCommand() *cobra.Command {
	client := initClient.InitClient()

	SecretCmd = &cobra.Command{
		Use:          "secrets [flags]",
		Short:        "list secrets",
		Example:      "kubectl secrets [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}

			err = ListSecretsWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return SecretCmd


}

func ListSecretsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	secretList, err := client.CoreV1().Secrets(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Secret名称", "Namespace", "Data个数", "Secret类型"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, secret := range secretList.Items {
		secretRow := []string{secret.Name, secret.Namespace, strconv.Itoa(len(secret.Data)), string(secret.Type)}
		if common.ShowLabels {
			secretRow = append(secretRow, common.LabelsMapToString(secret.Labels))
		}
		if common.ShowAnnotations {
			secretRow = append(secretRow, common.AnnotationsMapToString(secret.Annotations))
		}


		table.Append(secretRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}



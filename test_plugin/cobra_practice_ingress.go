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

var IngressCmd = &cobra.Command{}

func IngressCommand() *cobra.Command {
	client := initClient.InitClient()

	IngressCmd = &cobra.Command{
		Use:          "ingress [flags]",
		Short:        "list ingress",
		Example:      "kubectl ingress [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			err = ListIngressWithNamespace(client, ns)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return IngressCmd


}

func ListIngressWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	ingressList, err := client.NetworkingV1().Ingresses(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"Deployment名称", "Namespace", "IngressClassName", "Host"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	if common.ShowAnnotations {
		content = append(content, "Annotations")
	}

	table.SetHeader(content)


	for _, ingress := range ingressList.Items {
		ingressRow := []string{ingress.Name, ingress.Namespace, *ingress.Spec.IngressClassName, ingress.Spec.Rules[0].Host}
		if common.ShowLabels {
			ingressRow = append(ingressRow, common.LabelsMapToString(ingress.Labels))
		}
		if common.ShowAnnotations {
			ingressRow = append(ingressRow, common.AnnotationsMapToString(ingress.Annotations))
		}


		table.Append(ingressRow)
	}
	// 去掉表格线
	//table = common.TableSet(table)

	table.Render()

	return nil


}



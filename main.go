package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubectl_plugin_develop/initClient"
	"log"
)

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

	for _, pod := range podList.Items {
		fmt.Printf("pod Name:%s, pod Namespace:%s\n", pod.Name, pod.Namespace)
	}

	return nil
}
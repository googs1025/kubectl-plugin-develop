package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kubectl_plugin_develop/initClient"
	"log"
)

func main() {
	client := initClient.InitClient()

	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods ",
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
			list, err := client.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, pod := range list.Items {
				fmt.Println(pod.Name)
			}
			return nil
		},
	}

	MergeFlags(cmd)

	err := cmd.Execute()

	if err != nil {
		log.Fatalln(err)
	}
}

var cfgFlags *genericclioptions.ConfigFlags

func MergeFlags(cmd *cobra.Command) {
	cfgFlags = genericclioptions.NewConfigFlags(true)
	cfgFlags.AddFlags(cmd.Flags())
}

package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubectl_plugin_develop/initClient"
)

var ServiceCmd = &cobra.Command{}

func ServiceCommand() *cobra.Command {
	client := initClient.InitClient()

	ServiceCmd = &cobra.Command{
		Use:          "services [flags]",
		Short:        "list services",
		Example:      "kubectl services [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			list, err := client.CoreV1().Services(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, service := range list.Items {
				fmt.Println(service.Name)
			}
			return nil
		},
	}

	return ServiceCmd


}


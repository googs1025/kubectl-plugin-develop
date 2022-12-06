package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubectl_plugin_develop/initClient"
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
			list, err := client.AppsV1().Deployments(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, deployment := range list.Items {
				fmt.Println(deployment.Name)
			}
			return nil
		},
	}

	return DeploymentCmd


}

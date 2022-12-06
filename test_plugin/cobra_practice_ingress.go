package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubectl_plugin_develop/initClient"
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
			list, err := client.NetworkingV1().Ingresses(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, ingress := range list.Items {
				fmt.Println(ingress.Name)
			}
			return nil
		},
	}

	return IngressCmd


}


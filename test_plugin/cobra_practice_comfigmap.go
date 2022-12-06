package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubectl_plugin_develop/initClient"
)

var ConfigmapCmd = &cobra.Command{}

func ConfigmapCommand() *cobra.Command {
	client := initClient.InitClient()

	ConfigmapCmd = &cobra.Command{
		Use:          "configmaps [flags]",
		Short:        "list configmaps",
		Example:      "kubectl configmaps [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			list, err := client.CoreV1().ConfigMaps(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, configmap := range list.Items {
				fmt.Println(configmap.Name)
			}
			return nil
		},
	}

	return ConfigmapCmd


}

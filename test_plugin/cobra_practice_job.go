package main

import (
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubectl_plugin_develop/initClient"
	"context"
	"fmt"
)

var JobCmd = &cobra.Command{}

func JobCommand() *cobra.Command {
	client := initClient.InitClient()

	JobCmd = &cobra.Command{
		Use:          "jobs [flags]",
		Short:        "list jobs",
		Example:      "kubectl jobs [flags]",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			ns, err := c.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			if ns == "" {
				ns = "default"
			}
			list, err := client.BatchV1().Jobs(ns).List(context.Background(), v1.ListOptions{})
			if err != nil {
				return err
			}
			for _, job := range list.Items {
				fmt.Println(job.Name)
			}
			return nil
		},
	}

	return JobCmd


}


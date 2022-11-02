package main

import (
	"github.com/spf13/cobra"
	"kubectl_plugin_develop/common"
	"kubectl_plugin_develop/initClient"
	"kubectl_plugin_develop/pod"
	"log"
)

var _ = initClient.InitClient() // 需要先初始化

func main() {

	RunCmd(pod.PodRunCmd)
}



func RunCmd(f func(c *cobra.Command, args []string) error ) {
	cmd := &cobra.Command{
		Use:          "kubectl pods [flags]",
		Short:        "list pods ",
		Example:      "kubectl pods [flags]",
		SilenceUsage: true,
		RunE:f,
	}
	initClient.MergeFlags(cmd)
	//用来支持 是否 显示标签
	cmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl pods --show-lables")
	cmd.Flags().StringVar(&common.Labels,"labels","","kubectl pods --lables app=ngx or kubectl pods --lables=\"app=ngx,version=v1\"")
	cmd.Flags().StringVar(&common.Fields,"fields","","kubectl pods --fields=\"status.phase=Running\"")
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

/*
	-- fields 字段的限制
	metadata.name
	metadata.namespace
	spec.nodeName
	spec.restartPolicy
	spec.serviceAccountName
	status.phase
	status.podIP
	status.podIPs
	status.nominatedNodeName
 */
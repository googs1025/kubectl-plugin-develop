package main

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kubectl_plugin_develop/common"
	"log"


)

var cmdMetaData *common.CmdMetaData
func init() {
	cmdMetaData = &common.CmdMetaData{
		Use: "kubectl [flags]",
		Short: "kubectl list resources",
		Example: "kubectl [flags]",
	}

}

func main() {

	// 主命令
	mainCmd := &cobra.Command{
		Use:          cmdMetaData.Use,
		Short:        cmdMetaData.Short,
		Example:      cmdMetaData.Example,
		SilenceUsage: true,
	}

	// 各资源List命令
	podCmd := PodCommand()
	serviceCmd := ServiceCommand()
	deploymentCmd := DeploymentCommand()
	configmapCmd := ConfigmapCommand()
	ingressCmd := IngressCommand()
	jobCmd := JobCommand()
	// 注册
	MergeFlags(mainCmd, podCmd, serviceCmd, deploymentCmd, configmapCmd, ingressCmd, jobCmd)
	// 主command需要加入子command
	mainCmd.AddCommand(podCmd, serviceCmd, deploymentCmd, configmapCmd, ingressCmd, jobCmd)

	err := mainCmd.Execute() // 主命令执行

	if err != nil {
		log.Fatalln(err)
	}


}


var cfgFlags *genericclioptions.ConfigFlags


func MergeFlags(cmds ...*cobra.Command) {
	cfgFlags = genericclioptions.NewConfigFlags(true)
	for _, cmd := range cmds {
		cfgFlags.AddFlags(cmd.Flags())
	}

}
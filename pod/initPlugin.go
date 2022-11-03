package pod

import (
	"github.com/spf13/cobra"
	"kubectl_plugin_develop/common"
	"kubectl_plugin_develop/initClient"
	"log"
)
// 用来初始化kubectl pods插件用

var podCmdMetaData *common.CmdMetaData
func init() {
	podCmdMetaData = &common.CmdMetaData{
		Use: "kubectl pods [flags]",
		Short: "list pods ",
		Example: "kubectl pods [flags]",
	}

}

var _ = initClient.InitClient() // 需要先初始化

func RunCmd() {
	cmd := &cobra.Command{
		Use:          podCmdMetaData.Use,
		Short:        podCmdMetaData.Short,
		Example:      podCmdMetaData.Example,
		SilenceUsage: true,
	}
	// 合并主命令的参数
	initClient.MergeFlags(cmd, podListCmd, promptCmd)
	// 加入子命令参数

	// 用来支持输入命令行
	podListCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl pods --show-labels")
	podListCmd.Flags().StringVar(&common.Labels,"labels","","kubectl pods --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	podListCmd.Flags().StringVar(&common.Fields,"fields","","kubectl pods --fields=\"status.phase=Running\"")
	cmd.AddCommand(podListCmd, promptCmd)
	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}
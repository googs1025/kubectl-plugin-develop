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
		Use: "kubectl list [flags]",
		Short: "kubectl list resources",
		Example: "kubectl list [flags]",
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
	deploymentCmd := DeploymentCommand()
	statefulsetCmd := StatefulsetCommand()
	daemonsetCmd := DaemonsetCommand()
	jobCmd := JobCommand()
	cronjobCmd := CronjobCommand()

	serviceCmd := ServiceCommand()
	ingressCmd := IngressCommand()

	configmapCmd := ConfigmapCommand()
	secretCmd := SecretCommand()

	// 注册
	MergeFlags(mainCmd, podCmd, serviceCmd, deploymentCmd, configmapCmd, ingressCmd, jobCmd, statefulsetCmd, cronjobCmd, secretCmd, daemonsetCmd)
	// pods 支持标签、annotation返回，标签过滤、fields字段过滤
	podCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl pods --show-labels")
	podCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl pods --show-annotations")
	podCmd.Flags().StringVar(&common.Labels,"labels","","kubectl pods --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	podCmd.Flags().StringVar(&common.Fields,"fields","","kubectl pods --fields=\"status.phase=Running\"")
	// deployments
	deploymentCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl deployments --show-labels")
	deploymentCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl deployments --show-annotations")
	deploymentCmd.Flags().StringVar(&common.Labels,"labels","","kubectl deployments --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	deploymentCmd.Flags().StringVar(&common.Fields,"fields","","kubectl deployments --fields=\"status.phase=Running\"")
	// statefulsets
	statefulsetCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl statefulsets --show-labels")
	statefulsetCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl statefulsets --show-annotations")
	statefulsetCmd.Flags().StringVar(&common.Labels,"labels","","kubectl statefulsets --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	statefulsetCmd.Flags().StringVar(&common.Fields,"fields","","kubectl statefulsets --fields=\"status.phase=Running\"")
	// daemonsets
	daemonsetCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl daemonsets --show-labels")
	daemonsetCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl daemonsets --show-annotations")
	daemonsetCmd.Flags().StringVar(&common.Labels,"labels","","kubectl daemonsets --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	daemonsetCmd.Flags().StringVar(&common.Fields,"fields","","kubectl daemonsets --fields=\"status.phase=Running\"")
	// jobs
	jobCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl jobs --show-labels")
	jobCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl jobs --show-annotations")
	jobCmd.Flags().StringVar(&common.Labels,"labels","","kubectl jobs --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	jobCmd.Flags().StringVar(&common.Fields,"fields","","kubectl jobs --fields=\"status.phase=Running\"")
	// cronjobs
	cronjobCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl cronjobs --show-labels")
	cronjobCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl cronjobs --show-annotations")
	cronjobCmd.Flags().StringVar(&common.Labels,"labels","","kubectl cronjobs --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	cronjobCmd.Flags().StringVar(&common.Fields,"fields","","kubectl cronjobs --fields=\"status.phase=Running\"")
	// configmaps
	configmapCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl configmaps --show-labels")
	configmapCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl configmaps --show-annotations")
	configmapCmd.Flags().StringVar(&common.Labels,"labels","","kubectl configmaps --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	configmapCmd.Flags().StringVar(&common.Fields,"fields","","kubectl configmaps --fields=\"status.phase=Running\"")
	// secrets
	secretCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl secrets --show-labels")
	secretCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl secrets --show-annotations")
	secretCmd.Flags().StringVar(&common.Labels,"labels","","kubectl secrets --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	secretCmd.Flags().StringVar(&common.Fields,"fields","","kubectl secrets --fields=\"status.phase=Running\"")
	// services
	serviceCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl services --show-labels")
	serviceCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl services --show-annotations")
	serviceCmd.Flags().StringVar(&common.Labels,"labels","","kubectl services --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	serviceCmd.Flags().StringVar(&common.Fields,"fields","","kubectl services --fields=\"status.phase=Running\"")
	// ingress
	ingressCmd.Flags().BoolVar(&common.ShowLabels,"show-labels",false,"kubectl ingress --show-labels")
	ingressCmd.Flags().BoolVar(&common.ShowAnnotations,"show-annotations",false,"kubectl ingress --show-annotations")
	ingressCmd.Flags().StringVar(&common.Labels,"labels","","kubectl ingress --labels app=ngx or kubectl pods --labels=\"app=ngx,version=v1\"")
	ingressCmd.Flags().StringVar(&common.Fields,"fields","","kubectl ingress --fields=\"status.phase=Running\"")
	// 主command需要加入子command
	mainCmd.AddCommand(podCmd, serviceCmd, deploymentCmd, configmapCmd, ingressCmd, jobCmd, statefulsetCmd, cronjobCmd, daemonsetCmd, secretCmd)

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
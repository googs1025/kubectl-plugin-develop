package initClient

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"log"
)

var cfgFlags *genericclioptions.ConfigFlags

func InitClient() *kubernetes.Clientset {

	cfgFlags = genericclioptions.NewConfigFlags(true)
	// 可以注意一下里面的ClientConfig()与RawConfig()方法
	config, err := cfgFlags.ToRawKubeConfigLoader().ClientConfig()
	if err != nil {
		log.Fatalln(err)
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	return client

}

func MergeFlags(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cfgFlags.AddFlags(cmd.Flags())
	}

}

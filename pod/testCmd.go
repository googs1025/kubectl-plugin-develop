package pod

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"kubectl_plugin_develop/initClient"
	"sigs.k8s.io/yaml"
)

//做测试用， 无需纠结
var testCmd= &cobra.Command{
	Use:          "test",
	Short:        "test",
	Hidden:true,
	RunE: func(c *cobra.Command, args []string) error {
		client := initClient.InitClient()
		ns, err := c.Flags().GetString("namespace")
		if err != nil {
			return err
		}
		if ns == "" {
			ns="default"
		}

		pod, err := client.CoreV1().Pods(ns).Get(context.Background(),
			"hello-world-68fdbf5747-w789w",v1.GetOptions{})
		jsonStr, _ := json.Marshal(pod)

		ret := gjson.Get(string(jsonStr),"@this")


		m := make(map[string]interface{})
		err = yaml.Unmarshal([]byte(ret.Raw), &m)
		if err != nil {
			return nil
		}
		b, _ := yaml.Marshal(m)
		fmt.Println(string(b))
		return nil
	},

}

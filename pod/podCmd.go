package pod

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/kubernetes"
	"kubectl_plugin_develop/common"
	"kubectl_plugin_develop/initClient"
	"log"
	"os"
	"sigs.k8s.io/yaml"
)


//func PodRunCmd(cmd *cobra.Command,args []string) error {
//
//	client := initClient.InitClient()
//	ns, err := cmd.Flags().GetString("namespace")
//	if err != nil {
//		return err
//	}
//	if ns == "" {
//		ns = "default"
//	}
//	err = ListPodsWithNamespace(client, ns)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}


func ListPodsWithNamespace(client *kubernetes.Clientset, namespace string) error {
	ctx := context.Background()

	podList, err := client.CoreV1().Pods(namespace).List(ctx, v1.ListOptions{
		LabelSelector: common.Labels,
		FieldSelector: common.Fields,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// 表格化呈现
	table := tablewriter.NewWriter(os.Stdout)
	content := []string{"POD名称", "Namespace", "POD IP", "状态"}

	if common.ShowLabels {
		content = append(content, "标签")
	}
	table.SetHeader(content)

	//for _, pod := range podList.Items {
	//	fmt.Printf("pod Name:%s, pod Namespace:%s\n", pod.Name, pod.Namespace)
	//}

	for _, pod := range podList.Items {
		podRow := []string{pod.Name, pod.Namespace, pod.Status.PodIP, string(pod.Status.Phase)}

		if common.ShowLabels {
			podRow = append(podRow, common.LabelsMapToString(pod.Labels))
		}
		table.Append(podRow)
	}
	// 去掉表格线
	table = common.TableSet(table)

	table.Render()

	return nil
}

var podListCmd = &cobra.Command{
	Use:          "list",
	Short:        "list pods",
	Example:      "kubectl pods list [flags]",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := initClient.InitClient()
		ns, err := cmd.Flags().GetString("namespace")

		if err != nil {
			return err
		}
		if ns == "" {
			ns = "default"
		}
		err = ListPodsWithNamespace(client, ns)
		if err != nil {
			return err
		}
		return nil
	},

}

// 获取 pod详细
func getPodDetail(args []string, cmd *cobra.Command){
	// args如果没有写名称
	if len(args) == 0 {
		log.Println("podname is required")
		return
	}
	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		log.Println("error ns param")
		return
	}
	if ns == "" {
		ns = "default"
	}
	podName := args[0]
	pod, err := fact.Core().V1().Pods().Lister().
		Pods(ns).Get(podName)
	if err != nil {
		log.Println(err)
		return
	}
	b, err := yaml.Marshal(pod)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))

}


func getPodDetailByJSON(podName,path string,cmd *cobra.Command){
	ns,err:=cmd.Flags().GetString("namespace")
	if err!=nil{
		log.Println("error ns param")
		return }
	if ns==""{ns="default"}

	pod,err:=fact.Core().V1().Pods().Lister().
		Pods(ns).Get(podName)
	if err!=nil{
		log.Println(err)
		return
	}

	if path == PodEventType{ //代表 是取 POD事件
		eventsList, err := fact.Core().V1().Events().Lister().List(labels.Everything())
		if err != nil {
			log.Println(err)
			return
		}
		podEvents := []*corev1.Event{}
		for _, e := range eventsList {
			if e.InvolvedObject.UID == pod.UID {
				podEvents = append(podEvents, e)
			}
		}
		common.PrintEvent(podEvents)
		return
	}

	//获取日志
	if path==PodLogType{
		client := initClient.InitClient()
		req:=client.CoreV1().Pods(ns).GetLogs(pod.Name,&corev1.PodLogOptions{})	// PodLogOptions可以做设置
		ret:=req.Do(context.Background())
		b,err:= ret.Raw()
		if err!=nil{
			log.Println(err)
			return
		}
		fmt.Println(string(b))
		return
	}

	jsonStr,_:=json.Marshal(pod)
	ret:=gjson.Get(string(jsonStr),path)
	if !ret.Exists(){
		log.Println("无法找到对应的内容:"+path)
		return
	}
	if !ret.IsObject() && !ret.IsArray(){ //不是对象不是 数组，直接打印
		fmt.Println(ret.Raw)
		return
	}
	tempMap:=make(map[string]interface{})

	err=yaml.Unmarshal([]byte(ret.Raw),&tempMap)
	if err!=nil{
		log.Println(err)
		return
	}
	b,_:=yaml.Marshal(tempMap)
	fmt.Println(string(b))

}
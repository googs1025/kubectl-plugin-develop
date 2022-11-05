package pod

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/labels"
	"os"
	"regexp"
	"strings"
	"log"
)

func executorCmd(cmd *cobra.Command) func(in string) {
	return func(in string) {
		in = strings.TrimSpace(in)
		blocks := strings.Split(in, " ")
		args := make([]string,0)
		switch blocks[0] {
		case "exit":
			fmt.Println("Bye!")
			os.Exit(0)
		case "list":
			//if err := podListCmd.RunE(cmd,[]string{}); err!=nil {
			//	log.Fatalln(err)
			//}
			//InitCache() //初始化缓存
			//cacheCmd.ParseFlags(args)
			err := cacheCmd.RunE(cmd,args)
			if err != nil {
				log.Fatalln(err)
			}
		case "get":
			//getPodDetail(args, cmd)
			runtea(args, cmd)
		}
	}

}
var suggestions = []prompt.Suggest{
	// Command
	{"test", "this is test"},
	{"get", "获取POD详细"},
	{"exit", "退出交互式窗口"},
	{"list", "显示pod list列表"},
}

var promptCmd = &cobra.Command{
	Use:          "prompt",
	Short:        "prompt pods ",
	Example:      "kubectl pods prompt",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		InitCache() //初始化缓存
		p := prompt.New(
			executorCmd(c),
			completer,
			prompt.OptionPrefix(">>> "),
		)
		p.Run()
		return nil
	},

}

// getPodsList 单独拿出来用informer list
func getPodsList() (ret []prompt.Suggest) {
	podList, err := fact.Core().V1().Pods().Lister().
		Pods("default").List(labels.Everything())
	if err != nil {
		return
	}
	for _, pod := range podList{
		ret = append(ret, prompt.Suggest{
			Text: pod.Name,
			Description: "节点:"+pod.Spec.NodeName+" 状态:"+
				string(pod.Status.Phase)+" IP:"+pod.Status.PodIP,
		})
	}
	return
}
func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()	// 取到全部text
	if w == "" {
		return []prompt.Suggest{}
	}
	cmd, opt := parseCmd(in.TextBeforeCursor())	// 解析
	if cmd == "get" {
		return prompt.FilterHasPrefix(getPodsList(), opt, true)
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

func parseCmd(w string) (string, string) {
	// 使用正则表达，把多馀的空格去掉
	w = regexp.MustCompile("\\s+").ReplaceAllString(w," ")
	l := strings.Split(w," ")
	if len(l) >= 2 {
		return l[0], strings.Join(l[1:]," ")
	}
	return w, ""
}
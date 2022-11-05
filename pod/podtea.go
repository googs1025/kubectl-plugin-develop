package pod

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"log"
	"os"
)
type podjson struct {
	title string
	path string
}
type podmodel struct {
	items    []*podjson
	index   int
	cmd *cobra.Command
	podName string
}
func (m podmodel) Init() tea.Cmd {
	return nil
}
func (m podmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.index > 0 {
				m.index--
			}
		case "down":
			if m.index < len(m.items)-1 {
				m.index++
			}
		case "enter":
			getPodDetailByJSON(m.podName,m.items[m.index].path,m.cmd)
			return m,tea.Quit
		}
	}
	return m, nil
}

func (m podmodel) View() string {
	s := "按上下键选择要查看的内容\n\n"
	for i, item := range m.items {
		selected := " "
		if m.index == i {
			selected = "»"
		}
		s += fmt.Sprintf("%s %s\n", selected, item.title)
	}

	s += "\n按Q退出\n"
	return s
}

const (
	PodEventType = "__event__"
	PodLogType="__log__"

)

func runtea(args []string,cmd *cobra.Command)  {
	if len(args) == 0 {
		log.Println("podname is required")
		return
	}
	var podModel=podmodel{
		items:    []*podjson{},
		cmd: cmd,
		podName: args[0],
	}
	//v1.Pod{}
	podModel.items=append(podModel.items,
		&podjson{title:"元信息", path: "metadata"},
		&podjson{title: "标签", path: "metadata.labels"},
		&podjson{title: "注解", path: "metadata.annotations"},
		&podjson{title: "容器列表", path: "spec.containers"},
		&podjson{title: "上级依赖资源", path: "metadata.ownerReferences"},
		&podjson{title:"全部", path: "@this"},
		//下方的内容和 POD本身的yaml/json内容 无关
		&podjson{title:"*事件*", path: PodEventType},
		&podjson{title:"*日志*", path: PodLogType},
	)
	teaCmd := tea.NewProgram(podModel)
	if err := teaCmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		os.Exit(1)
	}
}

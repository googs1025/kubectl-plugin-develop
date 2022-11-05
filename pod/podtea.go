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

//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334

//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func runtea(args []string,cmd *cobra.Command)  {
	if len(args)==0{
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
		&podjson{title:"全部", path: "@this"},
	)
	teaCmd := tea.NewProgram(podModel)
	if err := teaCmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		os.Exit(1)
	}
}

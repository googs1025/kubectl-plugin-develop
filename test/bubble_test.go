package test

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"testing"
)

type model struct {
	items    []string
	index   int

}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// 键盘操作
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
			fmt.Println(m.items[m.index])
			return m,tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "欢迎来到程序员在囧途--k8s可视化系统\n\n"

	for i, item := range m.items {
		selected := " "
		if m.index == i {
			selected = "»"
		}
		s += fmt.Sprintf("%s %s\n", selected, item)
	}

	s += "\n按Q退出\n"
	return s
}
func TestBubble(t *testing.T) {

	var initModel = model{
		items:    []string{"我要看POD", "给我列出deployment", "我想看configmap"},
	}
	cmd := tea.NewProgram(initModel)
	if err := cmd.Start(); err != nil {
		fmt.Println("start failed:", err)
		os.Exit(1)
	}
}

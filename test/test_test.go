package test

import (
	"github.com/c-bata/go-prompt"
	"regexp"
	"strings"
	"fmt"
	"os"
	"testing"
)

func executor(in string) {
	in = strings.TrimSpace(in)

	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	}
}
var podSuggestions=[]prompt.Suggest{
	{"ngx-123", "ngx"},
	{"ngx-mygin", "mygin"},
	{"javapod-xxx", "javapod"},
}
//cmd optioins
func parseCmd(w string) (string,string){
	w=regexp.MustCompile("\\s+").ReplaceAllString(w," ")
	l:=strings.Split(w," ")
	if len(l)>=2{
		return l[0],strings.Join(l[1:]," ")
	}
	return w,""
}

var suggestions = []prompt.Suggest{
	// Command
	{"get", "get pod detail"},
	{"test", "this is test"},
	{"exit", "Exit prompt"},
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	cmd,opt:=parseCmd(in.TextBeforeCursor())
	if cmd=="get"{
		return prompt.FilterHasPrefix(podSuggestions,opt, true)
	}

	return prompt.FilterHasPrefix(suggestions, w, true)
}
//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
func TestCmd(t *testing.T) {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
	)
	p.Run()

	//本课程来自 程序员在囧途(www.jtthink.com) 咨询群：98514334
}

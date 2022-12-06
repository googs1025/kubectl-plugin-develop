package common

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	corev1 "k8s.io/api/core/v1"
	"os"
)

func LabelsMapToString(m map[string]string) (res string) {

	for k, v := range m {
		res += fmt.Sprintf("%s=%s,", k, v)
	}
	return
}

func AnnotationsMapToString(m map[string]string) (res string) {

	for k, v := range m {
		res += fmt.Sprintf("%s=%s, ", k, v)
	}
	return
}

func TableSet(table *tablewriter.Table) *tablewriter.Table {
	// 去掉表格线
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	return table
}


type CmdMetaData struct {
	Use string
	Short string
	Example string
}

var eventHeaders = []string{"事件类型", "REASON", "所属对象","消息"}

func PrintEvent(events []*corev1.Event){
	table := tablewriter.NewWriter(os.Stdout)
	//设置头
	table.SetHeader(eventHeaders)
	for _, e := range events {
		podRow := []string{e.Type,e.Reason,
			fmt.Sprintf("%s/%s",e.InvolvedObject.Kind,e.InvolvedObject.Name),e.Message}

		table.Append(podRow)
	}
	table = TableSet(table)
	table.Render()
}
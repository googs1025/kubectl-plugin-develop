package common

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
)

func MapToString(m map[string]string) (res string) {

	for k, v := range m {
		res += fmt.Sprintf("%s=%s \n", k, v)
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

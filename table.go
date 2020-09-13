package rld

import (
	"os"

	"github.com/guumaster/tablewriter"
)

func CreateTable(from string, to string, data []Section) {
	t := tablewriter.NewWriter(os.Stdout)

	for i := range data {
		if i > 0 {
			t.AddSeparator()
		}
		t.Append([]string{data[i].Title})
		t.AddSeparator()

		for j := range data[i].Entries {
			t.Append([]string{data[i].Entries[j][from], data[i].Entries[j][to]})
		}
	}

	t.SetAutoWrapText(true)
	t.SetReflowDuringAutoWrap(true)
	t.SetAutoMergeCells(true)
	t.SetCenterSeparator("")
	t.SetColumnSeparator("")
	t.SetHeaderLine(false)
	t.SetBorder(false)
	t.SetColMinWidth(0, 35)
	t.SetColMinWidth(1, 35)
	t.SetNoWhiteSpace(true)

	t.Render()
}

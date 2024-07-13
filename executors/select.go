package executors

import (
	data "polenta/data"
	sorter "polenta/sorter"
)

type SelectExecutor struct {
	statement string
}

func (s SelectExecutor) Execute() string {
	fields := "TBD"
	where := "TBD"
	orderBy := "TBD"
	rows := data.Rows{}
	selected := selectFrom(rows, fields, where, orderBy)
	return "Executed select statement. Selected " + string(rune(len(selected))) + " rows."
}

func selectFrom(rows data.Rows, _ string, _ string, orderBy string) data.Rows {
	return data.Rows(sorter.SortableRows(rows).Sort(orderBy))
}

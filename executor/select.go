package executor

import (
	data "polenta-db-go/data"
	sorter "polenta-db-go/sorter"
)

type SelectExecutor struct {
	statement string
}

func (s SelectExecutor) Execute() Response {
	fields := "TBD"
	where := "TBD"
	orderBy := "TBD"
	rows := data.Rows{}
	selected := selectFrom(rows, fields, where, orderBy)
	return Response{Message: "Executed select statement. Selected " + string(rune(len(selected))) + " rows."}
}

func selectFrom(rows data.Rows, _ string, _ string, orderBy string) data.Rows {
	return data.Rows(sorter.SortableRows(rows).Sort(orderBy))
}

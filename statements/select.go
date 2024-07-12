package statements

import (
	data "polenta/data"
	sorter "polenta/sorter"
)

type SelectStatement struct {
	statement string
}

func (s SelectStatement) Execute() string {
	return "Executed select statement"
}

func selectFrom(rows data.Rows, _ string, _ string, order string) data.Rows {
	return data.Rows(sorter.SortableRows(rows).Sort(order))
}

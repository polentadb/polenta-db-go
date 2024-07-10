package main

import (
	data "polenta/data"
	sorter "polenta/sorter"
)

func selectFrom(rows data.Rows, _ string, _ string, order string) data.Rows {
	return data.Rows(sorter.SortableRows(rows).Sort(order))
}

package main

import data "polenta/data"

func selectFrom(rows data.Rows, _ string, _ string, order string) data.Rows {
	return data.Rows(SortableRows(rows).Sort(order))
}

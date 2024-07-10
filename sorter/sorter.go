package sorter

import data "polenta/data"

type Sortable interface {
	Sort(criteria string) data.Rows
}

type SortableRows data.Rows

func (r SortableRows) Sort(criteria string) SortableRows {
	return selectionSort(r, criteria)
}

func exchange(rows SortableRows, i int, j int) {
	temp := rows[i]
	rows[i] = rows[j]
	rows[j] = temp
}

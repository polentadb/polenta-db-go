package main

type Sortable interface {
	sort(criteria string) Rows
}

func (r Rows) sort(criteria string) Rows {
	return selectionSort(r, criteria)
}

func exchange(rows Rows, i int, j int) {
	temp := rows[i]
	rows[i] = rows[j]
	rows[j] = temp
}

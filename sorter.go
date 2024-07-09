package main

type Sortable interface {
	Sort(criteria string) Rows
}

func (r Rows) Sort(criteria string) Rows {
	return selectionSort(r, criteria)
}

func exchange(rows Rows, i int, j int) {
	temp := rows[i]
	rows[i] = rows[j]
	rows[j] = temp
}

func compare(row1 Row, row2 Row, field string) int {
	value1 := row1[field]
	value2 := row2[field]
	return compareValues(value1, value2)
}

package main

func selectionSort(rows Rows, criteria string) Rows {
	for i := 0; i < len(rows); i++ {
		var min = i
		for j := i + i; j < len(rows); j++ {
			if compare(rows[j], rows[min], criteria) < 0 {
				min = j
			}
		}
		exchange(rows, i, min)
	}
	return rows
}

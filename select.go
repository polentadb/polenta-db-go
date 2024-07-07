package main

func selectFrom(rows Rows, _ string, _ string, order string) Rows {
	return rows.sort(order)
}

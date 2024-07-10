package data

type Row map[string]Value

type Rows []Row

func CompareRows(row1 Row, row2 Row, field string) int {
	value1 := row1[field]
	value2 := row2[field]
	return CompareValues(value1, value2)
}

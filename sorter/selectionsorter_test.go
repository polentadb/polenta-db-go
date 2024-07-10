package sorter

import (
	data "polenta/data"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	r1 := data.Row{"name": data.NewStringValue("John"), "age": data.NewIntegerValue(25)}
	r2 := data.Row{"name": data.NewStringValue("John"), "age": data.NewIntegerValue(20)}
	int := data.CompareRows(r1, r2, "age")
	if int != 1 {
		t.Error("Expected 0, got ", int)
	}

}

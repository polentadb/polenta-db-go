package main

import "testing"

func TestSelectionSort(t *testing.T) {

	r1 := Row{"name": NewStringValue("John"), "age": NewIntegerValue(25)}
	r2 := Row{"name": NewStringValue("John"), "age": NewIntegerValue(20)}
	int := compare(r1, r2, "age")
	if int != 1 {
		t.Error("Expected 0, got ", int)
	}

}

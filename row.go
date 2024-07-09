package main

type Value interface {
	//Compare(Value) int
}

type IntegerValue struct {
	value int
}

type StringValue struct {
	value string
}

func NewIntegerValue(v int) Value {
	var iv Value = IntegerValue{value: v}
	return iv
}

func NewStringValue(v string) Value {
	var iv Value = StringValue{value: v}
	return iv
}

type Row map[string]Value

type Rows []Row

func compareValues(v1 Value, v2 Value) int {
	switch v1.(type) {
	case IntegerValue:
		return v1.(IntegerValue).Compare(v2.(IntegerValue))
	case StringValue:
		return v2.(StringValue).Compare(v2.(StringValue))
	default:
		return 0
	}
}

func (v1 IntegerValue) Compare(v2 IntegerValue) int {
	if v1.value < v2.value {
		return -1
	} else if v1.value > v2.value {
		return 1
	} else {
		return 0
	}
}

func (v1 StringValue) Compare(v2 StringValue) int {
	if v1.value < v2.value {
		return -1
	} else if v1.value > v2.value {
		return 1
	} else {
		return 0
	}
}

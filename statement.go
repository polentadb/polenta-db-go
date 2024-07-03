package main

import (
	"fmt"
)

type ExecutableStatement interface {
	Execute() string
}

type Statement struct {
	statement string
}

func (s Statement) Execute() string {
	fmt.Println("Executing statement:", s.statement)
	return "Executed statement"
}

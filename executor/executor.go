package executor

import (
	"strings"
)

type Response struct {
	Success bool
	Message string
	Body    string
}

type Executor interface {
	Execute() Response
}

func Create(statement string) (Executor, *string) {
	if strings.HasPrefix(strings.ToUpper(statement), "CREATE") {
		return CreateExecutor{statement: statement}, nil
	}
	if strings.HasPrefix(strings.ToUpper(statement), "SELECT") {
		return SelectExecutor{statement: statement}, nil
	}
	if strings.HasPrefix(strings.ToUpper(statement), "INSERT") {
		return InsertExecutor{statement: statement}, nil
	}
	err := "Invalid statement"
	return nil, &err
}

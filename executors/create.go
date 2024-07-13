package executors

import (
	store "polenta/store"
)

type CreateExecutor struct {
	statement string
}

func (s CreateExecutor) Execute() string {
	execute(s.statement)
	return "Executed create statement"
}

func execute(statement string) {
	objectType := findObjectType(statement)
	objectName := findObjectName(statement)
	objectFields := findObjectFields(statement)
	store.Add(objectName, objectType, objectFields)
}

func findObjectName(_ string) string {
	return "TBD"
}

func findObjectType(_ string) string {
	return "bag"
}

func findObjectFields(_ string) map[string]string {
	var fields map[string]string
	return fields
}

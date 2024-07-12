package statements

import (
	"strings"
)

type Statement interface {
	Execute() string
}

func CreateExecutor(statement string) Statement {
	if strings.HasPrefix(strings.ToUpper(statement), "CREATE") {
		var createStatement CreateStatement = CreateStatement{statement: statement}
		return createStatement
	}
	if strings.HasPrefix(strings.ToUpper(statement), "SELECT") {
		var selectStatement SelectStatement = SelectStatement{statement: statement}
		return selectStatement
	}
	if strings.HasPrefix(strings.ToUpper(statement), "INSERT") {
		var insertStatement InsertStatement = InsertStatement{statement: statement}
		return insertStatement
	}
	var fallbackStatement FallbackStatement = FallbackStatement{statement: statement}
	return fallbackStatement
}

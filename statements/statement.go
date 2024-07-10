package statements

import (
	"strings"
)

type StatementExecutor interface {
	Execute() string
}

func CreateExecutor(statement string) StatementExecutor {
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

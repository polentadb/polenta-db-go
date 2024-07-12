package polenta

import statements "polenta/statements"

func Run(statement string) string {
	executor := statements.CreateExecutor(statement)
	if executor != nil {
		return executor.Execute()
	}
	return "Invalid statement"
}

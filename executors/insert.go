package executors

type InsertExecutor struct {
	statement string
}

func (s InsertExecutor) Execute() string {
	return "Executed insert statement"
}

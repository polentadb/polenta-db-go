package statements

type InsertStatement struct {
	statement string
}

func (s InsertStatement) Execute() string {
	return "Executed insert statement"
}

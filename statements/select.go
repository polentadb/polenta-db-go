package statements

type SelectStatement struct {
	statement string
}

func (s SelectStatement) Execute() string {
	return "Executed select statement"
}

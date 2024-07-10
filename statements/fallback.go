package statements

type FallbackStatement struct {
	statement string
}

func (s FallbackStatement) Execute() string {
	return "invalid statement: " + s.statement
}

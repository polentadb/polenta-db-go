package polenta

import executor "polenta/executor"

func Run(statement string) string {
	executor, err := executor.Create(statement)
	if err != nil {
		return *err
	}
	if executor != nil {
		return executor.Execute().Message
	}
	return "Invalid statement"
}

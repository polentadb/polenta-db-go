package polenta

import executors "polenta/executors"

func Run(statement string) string {
	executor, err := executors.Create(statement)
	if err != nil {
		return *err
	}
	if executor != nil {
		return executor.Execute()
	}
	return "Invalid statement"
}

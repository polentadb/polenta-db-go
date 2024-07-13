package polenta

import executor "polenta-db-go/executor"

func Run(statement string) executor.Response {
	exec, errCode, err := executor.Create(statement)
	if err != nil && errCode != nil {
		return executor.CreateErrorResponse(*errCode, *err)
	}
	if exec != nil {
		return exec.Execute()
	}
	return executor.CreateErrorResponse(2, "Internal error")
}

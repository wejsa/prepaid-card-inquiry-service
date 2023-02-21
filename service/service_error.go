package service

func GetErrorTasks(endpointMap dto.RequestEndpointMap) (errorTasks []dto.ErrorTask) {
	for _, v := range endpointMap {
		//log.Println("[" + v.TaskId() + "] ErrorMessage:" + v.ErrorMessage().GetCode())
		if v.ErrorMessage().GetCode() != "" {
			errorTask := dto.NewErrorTask(v.ComponentCode(), v.ErrorMessage())
			errorTasks = append(errorTasks, *errorTask)
		}
	}
	return
}

func GetResultByTasks(endpointMap dto.RequestEndpointMap) (doneTasks map[string]dto.ErrorTask, errorTasks map[string]dto.ErrorTask) {
	doneTasks = make(map[string]dto.ErrorTask)
	errorTasks = make(map[string]dto.ErrorTask)
	for _, v := range endpointMap {
		//log.Println("[" + v.TaskId() + "] ErrorMessage:" + v.ErrorMessage().GetCode())
		if v.ErrorMessage().GetCode() != "" {
			errorTask := dto.NewErrorTask(v.ComponentCode(), v.ErrorMessage())
			errorTasks[v.TaskId()] = *errorTask
		} else {
			doneTask := dto.NewErrorTask(v.ComponentCode(), v.ErrorMessage())
			doneTasks[v.TaskId()] = *doneTask
		}
	}
	return
}

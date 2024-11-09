package models

var taskMgr = TaskManager{}

func GetTaskManager() *TaskManager {
	return &taskMgr
}

func LoadFromFile(filePath string) error {
	taskMgr := GetTaskManager()
	return taskMgr.loadFromFile(filePath)
}

func SaveTaskManagerToFile(filePath string) error {
	taskMgr := GetTaskManager()
	return taskMgr.saveTaskManagerToFile(filePath)
}

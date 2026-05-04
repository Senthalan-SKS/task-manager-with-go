package service

import (
	"task-manager/model"
)

var tasks []model.Task

func GetTasks() []model.Task {
	return tasks
}

func AddTask(task model.Task) {
	tasks = append(tasks, task)
}
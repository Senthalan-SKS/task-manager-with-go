package handler

import (
	"encoding/json"
	"net/http"
	"task-manager/model"
	"task-manager/service"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := service.GetTasks()
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	json.NewDecoder(r.Body).Decode(&task)

	service.AddTask(task)

	json.NewEncoder(w).Encode(task)
}
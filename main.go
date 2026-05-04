package main

import (
	"log"
	"net/http"
	"task-manager/handler"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetTasks(w, r)
		} else if r.Method == http.MethodPost {
			handler.CreateTask(w, r)
		}
	})

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
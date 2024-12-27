package router

type Task struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Status bool   `json:"status"`
}

var AllTasks = make([]Task, 0)

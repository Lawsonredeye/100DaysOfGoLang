package router

type Task struct {
	Title       string `json:"title"`
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"status"`
}

var AllTasks = make([]Task, 0)

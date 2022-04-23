package domain

type Task struct {
	Done  bool
	Title string
}

type TaskList []Task

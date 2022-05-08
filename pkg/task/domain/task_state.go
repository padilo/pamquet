package domain

type TaskState struct {
	taskList TaskList
}

func InitState() TaskState {
	return TaskState{}
}

func (t TaskState) TaskList() TaskList {
	return t.taskList
}

func (t *TaskState) AddTask(title string) {
	t.taskList = append(t.taskList, Task{Title: title})
}

func (t *TaskState) GetTasksNames() []string {
	return collect(t.taskList, func(task Task) string {
		return task.Title
	})
}

func (t *TaskState) SetDone(selected int) {
	t.taskList[selected].Done = !t.taskList[selected].Done
}

func (t *TaskState) RemoveTask(selected int) {
	t.taskList = append(t.taskList[:selected], t.taskList[selected+1:]...)
}

func (t *TaskState) SetTitle(selected int, title string) {
	t.taskList[selected].Title = title
}

func (t *TaskState) SwitchTasks(i int, j int) {
	taskI := t.taskList[i]
	t.taskList[i] = t.taskList[j]
	t.taskList[j] = taskI
}

func collect[T any, U any](arrayItems []T, m func(T) U) []U {
	ret := make([]U, len(arrayItems))

	for i, item := range arrayItems {
		ret[i] = m(item)
	}

	return ret
}

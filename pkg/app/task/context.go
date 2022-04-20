package task

type Context struct {
	TaskList TaskList
}

func Init() Context {
	return Context{
		TaskList: TaskList{
			Task{Title: "test1"},
			Task{Title: "test2", Done: true},
			Task{Title: "test3"},
			Task{Title: "test4"},
		},
	}
}

func (c *Context) AddTask(title string) {
	c.TaskList = append(c.TaskList, Task{Title: title})
}

func (c *Context) GetTasksNames() []string {
	return collect[Task, string](c.TaskList, func(task Task) string {
		return task.Title
	})
}

func (c *Context) SetDone(selected int) {
	c.TaskList[selected].Done = !c.TaskList[selected].Done
}

func (c *Context) RemoveTask(selected int) {
	c.TaskList = append(c.TaskList[:selected], c.TaskList[selected+1:]...)
}

func (c *Context) SetTitle(selected int, title string) {
	c.TaskList[selected].Title = title
}

func (c *Context) SwitchTasks(i int, j int) {
	taskI := c.TaskList[i]
	c.TaskList[i] = c.TaskList[j]
	c.TaskList[j] = taskI
}

func collect[T any, U any](arrayItems []T, m func(T) U) []U {
	ret := make([]U, len(arrayItems))

	for i, item := range arrayItems {
		ret[i] = m(item)
	}

	return ret
}

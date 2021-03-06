package app

import (
	"github.com/padilo/pomaquet/pkg/task/domain"
)

type Context struct {
	TaskList domain.TaskList
}

func Init() Context {
	return Context{
		TaskList: domain.TaskList{},
	}
}

func (c *Context) AddTask(title string) {
	c.TaskList = append(c.TaskList, domain.Task{Title: title})
}

func (c *Context) GetTasksNames() []string {
	return collect[domain.Task, string](c.TaskList, func(task domain.Task) string {
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

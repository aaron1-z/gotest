package main

type Task struct {
	ID          int
	Description string
	Completed   bool
}

type TodoList struct {
	tasks []Task
}

func NewTodoList() *TodoList {
	return &TodoList{}
}

func (tl *TodoList) AddTask(description string) {
	task := Task{
		ID:          len(tl.tasks) + 1,
		Description: description,
		Completed:   false,
	}
	tl.tasks = append(tl.tasks, task)
}

func (tl *TodoList) ListTasks() []Task {
	return tl.tasks
}

func (tl *TodoList) CompleteTask(taskID int) {
	for i := range tl.tasks {
		if tl.tasks[i].ID == taskID {
			tl.tasks[i].Completed = true
			return
		}
	}
}

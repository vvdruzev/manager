package taskManager

import "log"

func NewTask(priority int) *Task {
	return &Task{Priority: priority, Observer: NewObserver()}
}

type Task struct {
	Priority int
	*Observer
}

func (t *Task) GetPriority() int {
	return t.Priority
}

func (t *Task) Execute() {
	err := t.StartObserve()
	if err != nil {
		log.Fatal("Error registering handlers")	}
	defer t.Close()
	t.Emit("create")
	t.Emit("exec")
	t.Emit("complete")
	t.Emit("error")

}
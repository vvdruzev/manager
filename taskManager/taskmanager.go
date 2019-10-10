package taskManager

import "log"

func NewTask(priority int) *Task {
	return &Task{Priority: priority, Observer: NewObserver()}
}

const (
	EventCreate string = "create"
	EventExec string = "exec"
	EventComplete string = "complete"
	EventError string = "error"
	)

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
		log.Fatal("Error registering listeners")
	}
	defer t.Close()
	t.Emit(EventCreate)
	t.Emit(EventExec)
	t.Emit(EventComplete)
	t.Emit(EventError)

}

package taskManager

import "testing"

func TestTask_Execute(t *testing.T) {
	var a,b,c,d int
	task := NewTask(1)
	task.AddListener(EventComplete, func(i interface{}) {
		a=1
	})
	task.AddListener(EventCreate, func(i interface{}) {
		b=1
	})
	task.AddListener(EventExec, func(i interface{}) {
		c=1
	})
	task.AddListener(EventError, func(i interface{}) {
		d=1
	})

	task.Execute()
	if a!=1 || b!=1 || c!=1 || d!=1 {
		t.Errorf(" got %d %d %d %d; want %d %d %d %d", a,d,c,d,1,1,1,1)
	}
}
package taskManager

import (
	"testing"
)

func TestInit0(t *testing.T) {
	h := new(PriorityQueue)
	task := NewTask(1)
	item := new(Item)
	item.Tasker = task
	for i := 20; i > 0; i-- {
		h.Add(item) // all elements are the same
	}
	Init(h)

	for i := 1; h.Len() > 0; i++ {
		x := h.Pop().(*Item)
		if x != item {
			t.Errorf("%v pop got %v; want %v", i, x, 0)
		}
	}
}

func TestInit1(t *testing.T) {
	h := new(PriorityQueue)
	task1 := NewTask(1)
	item := new(Item)
	item.Tasker = task1
	task2 := NewTask(2)
	item2 := new(Item)
	item2.Tasker = task2

	h.Add(item) // all elements are different
	h.Add(item2)
	Init(h)


	for i := h.Len(); h.Len()>0; i-- {
		x := h.Del().(*Item)
		if x.GetPriority() != i {
			t.Errorf("%v.th pop got %v; want %v", i, x.GetPriority(), i)
		}
	}
}


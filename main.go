// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)


type Tasker interface {
	Execute()
	GetPriority() int
}




// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	//priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
	Tasker
}

type task1 struct {
	priority int
}
func (t *task1) Execute()  {
	fmt.Printf("%T", t)
	//создание
	create()
	//Выполнение
	taskExec()
	//Завершение
	taskComplete()
	//Ошибка выполнения
	taskError()
}


func create() {
}

func taskExec()  {

}

func taskComplete()  {

}

func taskError()  {

}

func (t *task1) GetPriority() int  {
	return t.priority
}


type task2 struct {
	priority int
}

func (t *task2) Execute()  {
	fmt.Printf("%T", t)
}

func (t *task2) GetPriority() int  {
	return t.priority
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].GetPriority() > pq[j].GetPriority()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	t1:=task1{priority:1}
	t2 := task2{priority:2}
	//items := map[string]int{
	//	"t1": 3, "apple": 2, "pear": 4,
	//}
	////
	////// Create a priority queue, put the items in it, and
	////// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue,0)
	//i := 0
	//for value, priority := range items {
	//	pq[i] = &Item{
	//		value:    value,
	//		priority: priority,
	//		index:    i,
	//	}
	//	i++
	//}
	var t Tasker
	t = &t1
	pq = append(pq, &Item{value:"t1",index:0,Tasker:t})
	t = &t2
	pq = append(pq, &Item{value:"t2",index:1,Tasker:t})

	heap.Init(&pq)
	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "task13",
		Tasker:t,
	}
	heap.Push(&pq, item)
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		item.Execute()
		fmt.Printf("  %.2d:%s\n", item.GetPriority(), item.value)
	}
}

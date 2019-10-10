package taskManager

import "container/heap"

type Tasker interface {
	Execute()
	GetPriority() int
}


// An Item is something we manage in a priority queue.
type Item struct {
	Value    string // The value of the item; arbitrary.
	Index int // The index of the item in the heap.
	Tasker
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
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func Init(pq *PriorityQueue) {
	heap.Init(pq)
}

func (pq *PriorityQueue) Add(item *Item)  {
	heap.Push(pq,item)
}

func (pq *PriorityQueue) Del() interface{} {
	return heap.Pop(pq)
}
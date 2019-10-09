package main

import (
	tm "manager/taskManager"
	"fmt"
)

type task1 struct {
	name string
	*tm.Task
}
func main() {
	t2 := tm.NewTask(1)
	t1 := tm.NewTask(2)
	t3 := tm.NewTask(3)
	t4 := tm.NewTask(4)

	t5 := &task1{name:"type task1", Task:tm.NewTask(4)}

	t1.AddListener("create", func(i interface{}) {
		fmt.Println("create ", i.(string))
	})

	t1.AddListener("exec", func(i interface{}) {
		fmt.Println("exec", i.(string))
	})
	t1.AddListener("exec", func(i interface{}) {
		fmt.Println("exec1", i.(string))
	})

	t1.AddListener("exec", func(i interface{}) {
		fmt.Println("exec2", i.(string))
	})

	t2.AddListener("exec", func(i interface{}) {
		fmt.Println("exec", i.(string))
	})
	t2.AddListener("exec", func(i interface{}) {
		fmt.Println("exec1", i.(string))
	})

	t2.AddListener("exec", func(i interface{}) {
		fmt.Println("exec2", i.(string))
	})

	t3.AddListener("complete", func(i interface{}) {
		fmt.Println("complete", i.(string))
	})

	t4.AddListener("error", func(i interface{}) {
		fmt.Println("error", i.(string))
	})


	pq := make(tm.PriorityQueue,0)

	pq = append(pq, &tm.Item{Value:"t1",Index:0,Tasker:t1})

	pq = append(pq, &tm.Item{Value:"t2",Index:1,Tasker:t2})

	pq.Init()
	// Insert a new item and then modify its priority.
	pq.Add(&tm.Item{Value:"task4",Tasker:t4})
	pq.Add(&tm.Item{Value:"task3",Tasker:t3})
	pq.Add(&tm.Item{Value:"task5",Tasker:t5})

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := pq.Del().(*tm.Item)
		item.Execute()

		fmt.Printf("  %.2d:%s\n", item.GetPriority(), item.Value)
	}


}
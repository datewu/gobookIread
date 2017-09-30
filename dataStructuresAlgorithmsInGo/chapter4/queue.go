package main

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

func queueDemo() {
	q := queue.New()
	q.Enqueue(2)
	q.Enqueue(6)
	q.Enqueue(4)
	q.Enqueue(1)
	q.Enqueue(3)

	for q.Len() != 0 {
		fmt.Print(q.Dequeue())
	}

}

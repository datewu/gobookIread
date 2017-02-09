
package main

type myIntQueue struct {
	q []myInt
}

func newmyIntQueue() *myIntQueue {
	return &myIntQueue{
		q: []myInt{},
	}
}

func (q *myIntQueue) Insert(v myInt) {
	q.q = append(q.q, v)
}

func (q *myIntQueue) Remove()(res myInt) {
	if len(q.q) == 0 {
		panic("Oops")
	}
	res = q.q[0]
	q.q = q.q[1:]
	return
}

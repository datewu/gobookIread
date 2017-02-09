package main

type myTypeQueue struct {
	q []myType
}

func newMyTypeQueue() *myTypeQueue {
	return &myTypeQueue{
		q: []myType{},
	}
}

func (q *myTypeQueue) Insert(v myType) {
	q.q = append(q.q, v)
}

func (q *myTypeQueue) Remove(params) (res myType) {
	if len(q.q) == 0 {
		panic("Oops")
	}
	res = q.q[0]
	q.q = q.q[1:]
	return
}

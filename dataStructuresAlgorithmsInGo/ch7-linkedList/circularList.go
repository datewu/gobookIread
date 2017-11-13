package main

import "fmt"

type cList struct {
	tail  *cNode
	count int
}

type cNode struct {
	value int
	next  *cNode
}

func (l *cList) size() int {
	return l.count
}

func (l *cList) isEmpty() bool {
	return l.count == 0

}
func (l *cList) peek() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	return l.tail.next.value, true

}
func (l *cList) addHead(v int) {
	t := new(cNode)
	t.value = v
	if l.isEmpty() {
		l.tail = t
		t.next = t
	} else {
		t.next = l.tail.next
		l.tail.next = t
	}
	l.count++
}

func (l *cList) addTail(v int) {
	t := new(cNode)
	t.value = v
	if l.isEmpty() {
		l.tail = t
		t.next = t
	} else {
		t.next = l.tail.next
		l.tail.next = t
		l.tail = t
	}
	l.count++
}

func (l *cList) removeHead() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	v := l.tail.next.value
	if l.tail == l.tail.next {
		l.tail = nil
	} else {
		l.tail.next = l.tail.next.next
	}
	l.count--
	return v, true
}

func (l *cList) isPresent(v int) bool {
	t := l.tail
	for i := 0; i < l.count; i++ {
		if t.value == v {
			return true
		}
		t = t.next
	}
	return false
}

func (l *cList) print() {
	if l.isEmpty() {
		return
	}
	t := l.tail.next
	for t != l.tail {
		fmt.Print(t.value, " ")
		t = t.next
	}
	fmt.Println(t.value)
}

func (l *cList) freeList() {
	l.tail = nil
	l.count = 0
}

func (l *cList) removeNode(v int) bool {
	if l.isEmpty() {
		return false
	}
	prev, curr, head := l.tail, l.tail.next, l.tail.next
	if curr.value == v { // head
		if curr == curr.next { // single node case
			l.tail = nil
		} else {
			l.tail.next = l.tail.next.next
		}
		l.count--
		return true
	}
	prev = curr
	curr = curr.next

	for curr != head {
		if curr.value == v {
			if curr == l.tail {
				l.tail = prev
			}
			prev.next = curr.next
			l.count--
			return true
		}

		prev = curr
		curr = curr.next
	}
	return false
}

func (l *cList) copyReversed() *cList {
	cl := new(cList)
	curr := l.tail.next
	head := curr

	if curr != nil {
		cl.addHead(curr.value)
		curr = curr.next
	}

	for curr != head {
		cl.addHead(curr.value)
		curr = curr.next
	}
	return cl
}

func (l *cList) copyList() *cList {
	cl := new(cList)
	curr := l.tail.next
	head := curr

	if curr != nil {
		cl.addTail(curr.value)
		curr = curr.next
	}

	for curr != head {
		cl.addTail(curr.value)
		curr = curr.next
	}
	return cl
}

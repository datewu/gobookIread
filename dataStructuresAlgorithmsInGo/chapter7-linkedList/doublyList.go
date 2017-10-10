package main

import "fmt"

type dList struct {
	head, tail *dNode
	count      int
}

type dNode struct {
	value      int
	next, prev *dNode
}

func (l *dList) size() int {
	return l.count
}

func (l *dList) isEmpty() bool {
	return l.count == 0
}

func (l *dList) peek() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	return l.head.value, true
}

func (l *dList) addHead(v int) {
	newNode := &dNode{v, nil, nil}
	if l.count == 0 {
		l.tail = newNode
		l.head = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
	l.count++
}

func (l *dList) addTail(v int) {
	newNode := &dNode{v, nil, nil}
	if l.count == 0 {
		l.tail = newNode
		l.head = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
	l.count++
}

func (l *dList) removeHead() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	v := l.head.value
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}
	l.count--
	return v, true
}

func (l *dList) removeNode(v int) bool {
	curr := l.head
	if curr == nil {
		return false
	}
	if v == curr.value {
		curr = curr.next
		l.count--
		if curr != nil {
			l.head = curr
			l.head.prev = nil
		} else {
			l.tail = nil
		}
		return true
	}
	for curr.next != nil {
		if v == curr.next.value {
			curr.next = curr.next.next
			if curr.next == nil {
				l.tail = curr
			} else {
				curr.next.prev = curr
			}
			l.count--
			return true
		}
		curr = curr.next
	}
	return false
}

func (l *dList) isPresent(v int) bool {
	temp := l.head
	for temp != nil {
		if temp.value == v {
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *dList) freeList() {
	l.tail = nil
	l.head = nil
	l.count = 0
}

func (l *dList) print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

func (l *dList) reverseList() {
	curr := l.head
	var tempNode *dNode
	for curr != nil {
		tempNode = curr.next
		curr.next = curr.prev
		curr.prev = tempNode
		if curr.prev == nil {
			l.tail = l.head
			l.head = curr
			return
		}
		curr = curr.prev
	}
}

func (l *dList) copyListReversed(l2 *dList) {
	curr := l.head
	for curr != nil {
		l2.addHead(curr.value)
		curr = curr.next
	}
}

func (l *dList) copyList(l2 *dList) {
	curr := l.head
	for curr != nil {
		l2.addTail(curr.value)
		curr = curr.next
	}
}

func (l *dList) sortedInsert(v int) {
	temp := &dNode{v, nil, nil}
	curr := l.head

	if curr == nil { // first element
		l.head = temp
		l.tail = temp
		l.count++
		return
	}

	if l.head.value <= v { // at the begining
		temp.next = l.head
		l.head.prev = temp
		l.head = temp
		l.count++
		return
	}

	for curr.next != nil && curr.next.value > v {
		curr = curr.next
	}
	if curr.next == nil { // at the end
		l.tail = temp
		temp.prev = curr
		curr.next = temp
	} else { // all other
		temp.next = curr.next
		temp.prev = curr
		curr.next = temp
		temp.next.prev = temp
	}
	l.count++
}

// the list is already sorted
func (l *dList) removeDuplicate() {
	curr := l.head
	var deleteMe *dNode
	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			deleteMe = curr.next
			curr.next = deleteMe.next
			curr.next.prev = curr
			if deleteMe == l.tail {
				l.tail = curr
			}
		} else { // must have the else clause
			curr = curr.next
		}
	}
}

// func main() {
// 	list := new(dList)
// 	list.addHead(1)
// 	list.addHead(3)
// 	list.addHead(5)
// 	list.addHead(5)
// 	list.addHead(5)
// 	list.addHead(8)
// 	list.print()
// 	list.removeDuplicate()
// 	list.print()
// }

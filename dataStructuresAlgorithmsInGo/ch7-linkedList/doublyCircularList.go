package main

import "fmt"

type dcList struct {
	head, tail *dcNode
	count      int
}

type dcNode struct {
	value      int
	next, prev *dcNode
}

func (l *dcList) size() int {
	return l.count
}
func (l *dcList) isEmpty() bool {
	return l.count == 0
}
func (l *dcList) peekHead() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	return l.head.value, true
}

func (l *dcList) isPresent(v int) bool {
	t := l.head
	if l.head == nil {
		return false
	}
	for {
		if t.value == v {
			return true
		}
		t = t.next
		if t == l.head {
			break
		}
	}
	return false
}

func (l *dcList) freeList() {
	l.head = nil
	l.tail = nil
	l.count = 0
}

func (l *dcList) addHead(v int) {
	newNode := new(dcNode)
	newNode.value = v
	if l.isEmpty() {
		l.tail = newNode
		l.head = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = l.head
		newNode.prev = l.head.prev
		l.head.prev = newNode
		newNode.prev.next = newNode
		l.head = newNode
	}
	l.count++
}

func (l *dcList) addTail(v int) {
	newNode := new(dcNode)
	newNode.value = v
	if l.isEmpty() {
		l.head = newNode
		l.tail = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		newNode.next = l.tail.next
		newNode.prev = l.tail
		l.tail.next = newNode
		newNode.next.prev = newNode
		l.tail = newNode
	}
	l.count++
}

func (l *dcList) removeHead() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	v := l.head.value
	l.count--

	if l.isEmpty() {
		l.head = nil
		l.tail = nil
		return v, true
	}
	next := l.head.next
	next.prev = l.tail
	l.tail.next = next
	l.head = next
	return v, true
}

func (l *dcList) removeTail() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	v := l.tail.value
	l.count--

	if l.isEmpty() {
		l.head = nil
		l.tail = nil
		return v, true
	}
	prev := l.tail.prev
	prev.next = l.head
	l.head.prev = prev
	l.tail = prev
	return v, true
}

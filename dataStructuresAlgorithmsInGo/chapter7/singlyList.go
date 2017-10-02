package main

import "fmt"

type sList struct {
	head  *snode
	count int
}

type snode struct {
	value int
	next  *snode
}

func (l *sList) size() int {
	return l.count
}

func (l *sList) isEmpty() bool {
	return l.size() == 0
}

func (l *sList) addHead(value int) {
	l.head = &snode{value, l.head}
	l.count++
}

func (l *sList) addTail(value int) {
	curr := l.head
	newNode := &snode{value, nil}

	if curr == nil {
		l.head = newNode
		l.count++
		return
	}
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	l.count++
}

func (l *sList) print() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println("")
}

func (l *sList) sortedInsert(value int) {
	newNode := &snode{value, nil}
	curr := l.head

	if curr == nil || curr.value > value {
		newNode.next = l.head
		l.head = newNode
		l.count++
		return
	}
	for curr.next != nil && curr.next.value < value {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode
}

func (l *sList) isPresent(v int) bool {
	temp := l.head
	for temp != nil {
		if temp.value == v {
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *sList) removeHead() (int, bool) {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return 0, false
	}
	value := l.head.value
	l.head = l.head.next
	l.count--
	return value, true
}

func (l *sList) deleteNode(v int) bool {
	temp := l.head
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return false
	}
	if v == l.head.value {
		l.head = l.head.next
		l.count--
		return true
	}
	for temp.next != nil {
		if v == temp.next.value {
			temp.next = temp.next.next
			l.count--
			return true
		}
		temp = temp.next
	}
	return false
}

func (l *sList) deleteNodes(v int) bool {
	if l.isEmpty() {
		fmt.Println("Empty List Error")
		return false
	}

	currNode := l.head
	for currNode != nil && v == currNode.value {
		l.head = currNode.next
		currNode = l.head
		l.count--
	}
	for currNode != nil {
		nextNode := currNode.next
		if nextNode != nil && v == nextNode.value {
			currNode.next = nextNode.next
			l.count--
		} else {
			currNode = nextNode
		}
	}
	return false
}

func (l *sList) freeList() {
	l.head = nil
	l.count = 0
}

func (l *sList) reverse() {
	curr := l.head
	var prev, next *snode

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
}

func (l *sList) reverseWithRecurse() {
	var r func(*snode, *snode) *snode
	r = func(currentNode, prevNode *snode) *snode {
		var ret *snode
		if currentNode == nil {
			return nil
		}
		if currentNode.next == nil {
			currentNode.next = prevNode
			return currentNode
		}
		ret = r(currentNode.next, currentNode)
		currentNode.next = prevNode
		return ret
	}
	l.head = r(l.head, nil)
}

// the list is sorted
func (l *sList) removeDuplicate() {
	curr := l.head
	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}

func (l *sList) copyListReversed() *sList {
	var temp1, temp2 *snode
	curr := l.head

	for curr != nil {
		temp2 = &snode{curr.value, temp1}
		temp1 = temp2
		curr = curr.next
	}
	list := new(sList)
	list.head = temp1
	return list
}

func (l *sList) copyList() *sList {
	var headNode, tailNode, tempNode *snode
	curr := l.head
	if curr == nil {
		list := new(sList)
		list.head = nil
		return list
	}

	headNode = &snode{curr.value, nil}
	tailNode = headNode
	curr = curr.next
	for curr != nil {
		tempNode = &snode{curr.value, nil}
		tailNode.next = tempNode
		tailNode = tempNode
		curr = curr.next
	}

	list := new(sList)
	list.head = headNode
	return list
}

func (l *sList) compareList(ll *sList) bool {
	var r func(*snode, *snode) bool

	r = func(head1, head2 *snode) bool {
		if head1 == nil && head2 == nil {
			return true
		} else if head1 == nil || head2 == nil || head1.value != head2.value {
			return false
		} else {
			return r(head1.next, head2.next)
		}
	}

	return r(l.head, ll.head)
}

func (l *sList) nthFromBeigning(index int) (int, bool) {
	if index > l.size() || index < 1 {
		fmt.Println("too few nodes")
		return 0, false
	}
	curr := l.head
	for count := 0; curr != nil && count < index-1; count++ {
		curr = curr.next
	}
	return curr.value, true
}

func (l *sList) nthFromEnd(index int) (int, bool) {
	if index > l.size() || index < 1 {
		fmt.Println("too few nodes")
		return 0, false
	}
	size := l.size()
	i := size - index + 1
	return l.nthFromBeigning(i)
}

func (l *sList) nthFromEnd2(index int) (int, bool) {
	count := 1
	curr := l.head
	forward := curr

	for ; forward != nil && count <= index; count++ {
		forward = forward.next
	}
	if forward == nil {
		fmt.Println("too few nodes")
		return 0, false
	}
	for forward != nil {
		forward = forward.next
		curr = curr.next
	}
	return curr.value, true
}

// slow reference and fast reference (SPFP)
func (l *sList) loopDetect() bool {
	slowPtr, fastPtr := l.head, l.head

	for fastPtr.next != nil && fastPtr.next.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("loop found")
			return true
		}
	}
	fmt.Println("loop not found")
	return false
}

func (l *sList) reverseListLoopDetect() bool {
	originHead := l.head
	l.reverse()
	if originHead == l.head {
		l.reverse()
		fmt.Println("loop found")
		return true
	}
	l.reverse()
	fmt.Println("loop not found")
	return false
}

func (l *sList) loopTypeDetect() int {
	slowPtr, fastPtr := l.head, l.head

	for fastPtr.next != nil && fastPtr.next.next != nil {
		if l.head == fastPtr.next || l.head == fastPtr.next.next {
			fmt.Println("circular list loop found")
			return 2
		}

		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			fmt.Println("loop found")
			return 1
		}
	}
	fmt.Println("loop not found")
	return 0
}

func (l *sList) removeLoop() {
	intersectionPoint := func(list *sList) *snode {
		slowPtr := list.head
		fastPtr := slowPtr
		for fastPtr.next != nil && fastPtr.next.next != nil {
			slowPtr = slowPtr.next
			fastPtr = fastPtr.next.next
			if slowPtr == fastPtr {
				return fastPtr
			}
		}
		return nil
	}
	loopPoint := intersectionPoint(l)
	if loopPoint == nil {
		return
	}
	firstPtr := l.head
	if loopPoint == firstPtr {
		for firstPtr.next != l.head {
			firstPtr = firstPtr.next
		}
		firstPtr.next = nil
		return
	}
	secondPtr := loopPoint
	for firstPtr.next != secondPtr.next {
		firstPtr = firstPtr.next
		secondPtr = secondPtr.next
	}
	secondPtr.next = nil

}

// “Given two linked list which meet at some point find that intersection point.”
//
// 摘录来自: Hemant Jain. “Data Structures & Algorithms In Go”。 iBooks.
func (l *sList) findIntersection(head1, head2 *snode) *snode {
	var l1, l2, diff int
	temp1, temp2 := head1, head2
	for temp1 != nil {
		l1++
		temp1 = temp1.next
	}
	for temp2 != nil {
		l2++
		temp2 = temp2.next
	}
	if l1 < l2 {
		t := head1
		head1 = head2
		head2 = t
		diff = l2 - l1
	} else {
		diff = l1 - l2
	}

	for ; diff > 0; diff++ {
		head1 = head1.next
	}
	for head1 != head2 {
		head1 = head1.next
		head2 = head2.next
	}
	return head1
}

// func main() {
// 	list := &sList{}
// 	list.addHead(2)
// 	list.addHead(5)
// 	list.addHead(8)
// 	list.addHead(50)
// 	list.addHead(52)
// 	list.addHead(55)
// 	l := list.copyList()
// 	l.addTail(88)
// 	list.print()
// 	l.print()
// 	fmt.Println("compare self", list.compareList(list))
// 	fmt.Println("compare diff", list.compareList(l))
// }

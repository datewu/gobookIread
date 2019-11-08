package main

import (
	"fmt"
	"container/list"
)

func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(41)
	intList.PushBack(13)
	intList.PushBack(31)
	intList.PushBack(29)
	intList.PushBack(49)

	for element := intList.Front(); element != nil ;element = element.Next() {
		fmt.Println(element, element.Value.(int))
	}
}
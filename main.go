package main

import (
	"dataStructuresInGo/linkedList"
	"dataStructuresInGo/linkedListUtils"
	"fmt"
)

func main() {
	var ll = &linkedList.LinkedList{}
	for i := 0; i < 10; i++ {
		ll.AddTail((i + 1) * (i + 2))

	}

	linkedList.Traverse(ll.GetHead())

	fmt.Println("This was Linked List")

	newHead := linkedListUtils.RemoveNthNode(ll.GetHead(), 3)
	linkedList.Traverse(newHead)
	fmt.Println("updated Linked List")

}

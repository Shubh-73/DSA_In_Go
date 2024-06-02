package main

import (
	"dataStructuresInGo/linkedList"
	"dataStructuresInGo/linkedListUtils"
	"fmt"
)

func main() {
	var ll = &linkedList.NewLinkedList{}
	for i := 0; i < 10; i++ {
		ll.AddTail((i + 1) * (i + 2))

	}
	fmt.Println("Original Linked List ")
	linkedList.Traverse(ll.GetHead())
	fmt.Println("============================================")

	fmt.Println("======finding the middle of linked list ===========")
	midNodeValue, _ := linkedListUtils.FindMiddle(ll.GetHead())
	linkedList.Traverse(ll.GetHead())
	fmt.Println("Middle Node of Linked List -> ", midNodeValue)
	fmt.Println("============================================")
	fmt.Println("======deleting the middle of linked list ==========")
	newHead := linkedListUtils.DeleteMidNode(ll.GetHead())
	linkedList.Traverse(newHead)
}

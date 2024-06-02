package linkedListUtils

import (
	"dataStructuresInGo/linkedList"
)

func RemoveNthNode(head *linkedList.Node, n int) *linkedList.Node {
	if head == nil || n <= 0 {
		return head
	}

	if n == 1 {
		return head.Next
	}

	current := head

	for i := 0; current != nil && i < n-1; i++ {
		current = current.Next
	}
	if current != nil && current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}

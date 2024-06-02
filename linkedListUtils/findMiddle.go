package linkedListUtils

import (
	"dataStructuresInGo/linkedList"
)

func FindMiddle(head *linkedList.Node) (int, bool) {
	if head == nil {
		return 0, false
	}

	slowPointer := head
	fastPointer := head

	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
	}

	return slowPointer.Value, true
}

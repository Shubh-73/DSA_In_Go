package linkedListUtils

import "dataStructuresInGo/linkedList"

func DeleteMidNode(head *linkedList.Node) *linkedList.Node {

	if head == nil || head.Next == nil {
		return nil
	}

	slowPointer := head
	fastPointer := head
	previousToSlowPointer := head
	for fastPointer.Next != nil && fastPointer.Next.Next != nil {
		fastPointer = fastPointer.Next.Next
		previousToSlowPointer = slowPointer
		slowPointer = slowPointer.Next

	}

	if previousToSlowPointer != nil && slowPointer != nil {
		previousToSlowPointer.Next = previousToSlowPointer.Next.Next
	}

	return head
}

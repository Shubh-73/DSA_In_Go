package stacks

import (
	"dataStructuresInGo/linkedList"
)

type Stack struct {
	list linkedList.NewLinkedList
}

func NewStack() *Stack {
	return &Stack{
		list: linkedList.NewLinkedList{},
	}
}
func (stack *Stack) Push(value int) {
	stack.list.AddHead(value)
}

func (stack *Stack) Pop() (int, bool) {
	return stack.list.RemoveHead()
}

func (stack *Stack) Peek() (int, bool) {
	head := stack.list.GetHead()
	if head == nil {
		return 0, false
	}
	return head.Value, true
}

func (stack *Stack) IsEmpty() bool {

	return stack.list.GetHead() == nil
}

func (stack *Stack) Size() int {
	size := 0
	current := stack.list.GetHead()
	for current != nil {
		size++
	}
	return size
}

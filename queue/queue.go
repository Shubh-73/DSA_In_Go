package queue

import (
	"dataStructuresInGo/linkedList"
)

type Queue struct {
	list linkedList.NewLinkedList
}

func NewQueue() *Queue {
	return &Queue{
		list: linkedList.NewLinkedList{},
	}
}
func (this *Queue) Enqueue(value int) {
	this.list.AddTail(value)
}
func (this *Queue) Dequeue() (int, bool) {
	return this.list.RemoveHead()
}

func (this *Queue) IsEmpty() bool {

	return this.list.GetHead() == nil
}
func (this *Queue) GetSize() int {
	size := 0
	current := this.list.GetHead()
	for current != nil {
		size++
	}
	return size
}

func (this *Queue) Peek() (int, bool) {
	head := this.list.GetHead()
	if head == nil {
		return 0, false
	}
	return head.Value, true
}

package linkedList

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type NewLinkedList struct {
	Head *Node
	Tail *Node
}

func (this *NewLinkedList) AddHead(value int) *Node {
	newNode := &Node{Value: value, Next: this.Head}
	if this.Head == nil {
		this.Tail = newNode
	}
	this.Head = newNode
	return newNode
}

func (this *NewLinkedList) AddTail(value int) *Node {
	newNode := &Node{Value: value}
	if this.Tail != nil {
		this.Tail.Next = newNode
	} else {
		this.Head = newNode
	}
	this.Tail = newNode
	return this.Head

}

func (this *NewLinkedList) RemoveHead() (int, bool) {
	if this.Head == nil {
		return 0, false
	}
	currentNode := this.Head
	newHead := currentNode.Next
	currentNode.Next = nil
	return newHead.Value, true
}

func (this *NewLinkedList) GetHead() *Node {
	return this.Head
}

func Traverse(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Print(currentNode.Value, " -> ")
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

package main

import "fmt"

func main() {
	l := NewLinkedList()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.Insert(4)
	l.Insert(5)
	l.Print()
	l.Delete(3)
	l.Print()
}

type CustomizedData int

type Node struct {
	Next *Node
	Data CustomizedData
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(data CustomizedData) {
	n := &Node{
		Next: nil,
		Data: data,
	}

	if l.Head == nil {
		l.Head = n
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = n

}

func (l *LinkedList) Print() {
	curr := l.Head
	for curr != nil {
		if curr.Next!=nil{
			fmt.Printf(" %d -> ",curr.Data)
			curr = curr.Next
		}else{
			fmt.Print(curr.Data)
			curr = curr.Next
		}
		
	}
	println()
}

func (l *LinkedList) Delete(data CustomizedData){
	prev:=l.Head
	curr:=l.Head

	for curr!=nil{
		if curr.Data==data{
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
	}
}
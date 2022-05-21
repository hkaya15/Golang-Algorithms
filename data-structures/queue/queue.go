package main

import "fmt"

func main(){
	q:=newQueue()
	q.Enqueue(2)
	q.Enqueue(4)
	q.Enqueue(6)
	q.Enqueue(8)
	println(q.Dequeue())
	println(len(q.list))
	fmt.Println(q.list)
}

type CustomizedData int

type Queue struct{
	list []CustomizedData
}

func newQueue() *Queue{
	return &Queue{list: []CustomizedData{}}
}

func (q *Queue) Enqueue(value CustomizedData){
	q.list=append(q.list,value )
}

func (q *Queue) Dequeue()CustomizedData{
	first:=q.list[0]
	q.list=q.list[1:]
	return first
}
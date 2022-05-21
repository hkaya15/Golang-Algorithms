package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error! :", r)
		}
	}()
	s := newStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for i := 0; i < 2; i++ {
		fmt.Println(s.Pop())
	}
	fmt.Println(s.list)

}

type CustomizedData int

type Stack struct {
	cap   uint64
	depth uint64
	list  []CustomizedData
}

func newStack(cap uint) *Stack {
	return &Stack{
		cap:   uint64(cap),
		depth: 0,
		list:  make([]CustomizedData, cap),
	}
}

func (s *Stack) Push(value CustomizedData) {
	if s.depth == s.cap {
		panic("Out of capacity")
	}
	s.list[s.depth] = value
	s.depth++
}

func (s *Stack) Pop() CustomizedData {
	if s.depth < 1 {
		return -1
	} else {
		s.depth--
		last := s.list[s.depth]
		s.list = s.list[:len(s.list)-1]
		return last
	}
}

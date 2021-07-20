package main

import "fmt"

type List struct {
	Value int
	Next *List
}

func main() {
	l1 := new(List)
	l1.Value = 1
	l2 := new(List)
	l2.Value = 2
	l3 := new(List)
	l3.Value = 3
	l1.AppendList(l2)
	l1.InsertList(l3)
	PrintList(l1)
}

func PrintList(l *List) {
	for l != nil {
		fmt.Println(l.Value)
		l = l.Next
	}
}

func (l *List)InsertList(instance *List) {
}

func (l *List)AppendList(instance *List) {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = instance
}

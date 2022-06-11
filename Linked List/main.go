package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Data	T
	Next	*Node[T]
}

type LinkedList[T comparable] struct {
	Head	*Node[T]
}

func (l *LinkedList[T]) size() int {
	count := 0
	temp := l.Head
	for ;; {
		if temp.Next == nil {
			break
		}
		count++
		temp = temp.Next
	}
	return count
}

func (l *LinkedList[T]) insert(data T, index int) {
	newNode := &Node[T]{Data: data}
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
	} else {
		temp := l.Head
		for i := 0; i < index - 1; i++ {
			temp = temp.Next
		}
		newNode.Next = temp.Next
		temp.Next = newNode
	}
}

func (l *LinkedList[T]) remove(index int) {
	if l.Head == nil || index < 0 || index > l.size() {
		return;
	}
	if index == 0 {
		l.Head = l.Head.Next
	} else {
		temp := l.Head
		for i := 0; i < index - 1; i++ {
			temp = temp.Next
		}
		temp.Next = temp.Next.Next
	}
}

func (l *LinkedList[T]) reverse() {
	var prev, curr, next *Node[T]
	curr = l.Head
	for ;; {
		if curr == nil {
			break
		}
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

func (l *LinkedList[T]) print() {
	temp := l.Head
	for ;; {
		if temp.Next == nil {
			break
		}
		fmt.Print(temp.Data, " -> ")
		temp = temp.Next
	}
	fmt.Println(temp.Data)
}

func main() {
	ll := LinkedList[int]{Head: nil}
	ll.print()
}


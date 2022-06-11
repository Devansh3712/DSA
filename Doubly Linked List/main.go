package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Data T
	Prev *Node[T]
	Next *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	Head *Node[T]
}

func (d *DoublyLinkedList[T]) Size() int {
	count := 0
	temp := d.Head
	for {
		if temp == nil {
			break
		}
		count++
		temp = temp.Next
	}
	return count
}

func (d *DoublyLinkedList[T]) Insert(data T, index int) {
	newNode := &Node[T]{Data: data}
	if index == 0 {
		newNode.Prev = nil
		newNode.Next = d.Head
		if d.Head != nil {
			d.Head.Prev = newNode
		}
		d.Head = newNode
	} else {
		temp := d.Head
		for i := 0; i < index-1; i++ {
			temp = temp.Next
		}
		newNode.Prev = temp
		if index == d.Size() {
			newNode.Next = nil
		} else {
			newNode.Next = temp.Next
			temp.Next.Prev = newNode
		}
		temp.Next = newNode
	}
}

func (d *DoublyLinkedList[T]) Remove(index int) {
	if d.Head == nil || index < 0 || index > d.Size()-1 {
		return
	}
	if index == 0 {
		d.Head = d.Head.Next
		d.Head.Prev = nil
	} else {
		temp := d.Head
		for i := 0; i < index-1; i++ {
			temp = temp.Next
		}
		if index == d.Size()-1 {
			temp.Next = nil
		} else {
			temp.Next.Next.Prev = temp
			temp.Next = temp.Next.Next
		}
	}
}

func (d *DoublyLinkedList[T]) Reverse() {
	if d.Head == nil || d.Head.Next == nil {
		return
	}
	var temp, curr *Node[T]
	curr = d.Head
	for {
		if curr == nil {
			break
		}
		temp = curr.Prev
		curr.Prev = curr.Next
		curr.Next = temp
		curr = curr.Prev
	}
	if temp != nil {
		d.Head = temp.Prev
	}
}

func (d *DoublyLinkedList[T]) Print() {
	temp := d.Head
	for {
		if temp.Next == nil {
			break
		}
		fmt.Print(temp.Data, " <=> ")
		temp = temp.Next
	}
	fmt.Println(temp.Data)
}

func main() {
	dll := DoublyLinkedList[int]{Head: nil}
	dll.Print()
}

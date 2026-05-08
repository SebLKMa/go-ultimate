package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(newData interface{}) {
	// Insert always insert at end
	newNode := &Node{data: newData, next: nil}
	if l.head == nil {
		// empty list, just set data to head
		l.head = newNode
	} else {
		// iterate to last node, set data to last node
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *LinkedList) Delete(index int) {
	if l.head == nil {
		return
	}

	// To delete the first node
	if index == 0 {
		if l.head != nil {
			l.head = l.head.next
		}
		return
	}

	// To set previous node to point to next node
	pos := 0
	var previous *Node
	current := l.head
	for current.next != nil {
		if index == pos {
			previous.next = current.next
			break
		}
		previous = current
		current = current.next
		pos++
	}

	// To delete the last node
	if index == pos {
		previous.next = current.next
	}
}

func (l *LinkedList) Print() {
	current := l.head
	for current.next != nil {
		fmt.Printf("-> %v", current.data)
		current = current.next
	}
	fmt.Printf("-> %v", current.data) // the last node
	fmt.Println()
}

func (l *LinkedList) ToCircular() {
	// Setting the last node to point to the first node
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = l.head
}

func (l *LinkedList) PrintCircular() {
	current := l.head
	for current.next != nil {
		if current.next == l.head {
			// already reached starting point
			fmt.Printf("-> %v->[%v]", current.data, current.next.data)
			break
		}
		fmt.Printf("-> %v", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	ll := LinkedList{}
	for i := 0; i <= 5; i++ {
		ll.Insert(rand.Intn(100)) // for random int less than 100
	}
	ll.Print()

	ll2 := LinkedList{}
	for i := 0; i <= 5; i++ {
		ll2.Insert(i) // just int
	}
	ll2.Print()
	ll2.Delete(3)
	ll2.Print()

	ll2.ToCircular()
	ll2.PrintCircular()
}

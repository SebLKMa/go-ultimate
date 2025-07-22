package main

import "fmt"

const QSize = 1024

// Queue - first in first out
type Queue struct {
	q    []string
	qptr int // Always point to the next free slot. Top of the stack is stack[sp-1]
	// Incremented and decremented as the queue grows or shrinks.
}

func New() *Queue {
	return &Queue{q: make([]string, QSize)}
}

func (q *Queue) Enqueue(data string) error {
	if q.qptr >= QSize {
		return fmt.Errorf("queue exceeded")
	}

	q.q[q.qptr] = data
	q.qptr++
	return nil
}

func (q *Queue) Dequeue() string {
	// Get the top element, then move the rest up
	data := q.q[0]
	max := q.qptr
	q.qptr--
	for i := 1; i <= max; i++ {
		q.q[i-1] = q.q[i]
	}

	return data
}

func (q *Queue) Size() int {
	return q.qptr
}

func (q *Queue) Print() {
	for i, data := range q.q {
		if i >= q.qptr {
			break
		}
		fmt.Println(data)
	}
}

func main() {
	myQ := New()
	myQ.Enqueue("GoodBye")
	myQ.Enqueue("Hello")
	myQ.Enqueue("World")
	myQ.Print()
	fmt.Println("Queue size: ", myQ.Size())

	fmt.Println("Dequeue: ", myQ.Dequeue())
	fmt.Println("Queue size: ", myQ.Size())
	fmt.Println("Dequeue: ", myQ.Dequeue())
	fmt.Println("Queue size: ", myQ.Size())
	fmt.Println("Dequeue: ", myQ.Dequeue())
	fmt.Println("Queue size: ", myQ.Size())

}

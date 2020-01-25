package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	queue := NewCircularQueue(5)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	queue.Enqueue(5)
	queue.Enqueue(6)
	fmt.Println(queue)
	queue.Enqueue(7)
	fmt.Println(queue)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	queue.Enqueue(8)
	queue.Enqueue(9)
	queue.Enqueue(10)
	fmt.Println(queue)
}

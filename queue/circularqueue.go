package queue

import "fmt"

type CircularQueue struct {
	items []interface{}
	head  int
	tail  int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]interface{}, capacity+1),
		head:  0,
		tail:  0,
	}
}

func (c *CircularQueue) Enqueue(item interface{}) bool {
	if (c.tail+1)%len(c.items) == c.head {
		return false
	}
	c.items[c.tail] = item
	c.tail = (c.tail + 1) % len(c.items)
	return true
}

func (c *CircularQueue) Dequeue() interface{} {
	if c.head == c.tail {
		return nil
	}
	ret := c.items[c.head]
	c.head = (c.head + 1) % len(c.items)
	return ret
}

func (c *CircularQueue) String() string {
	var str string
	for i := c.head; ; {
		str += fmt.Sprintf("%+v", c.items[i])
		i = (i + 1) % len(c.items)
		if i != c.tail {
			str += " "
		} else {
			break
		}
	}
	return fmt.Sprintf("[%s]", str)
}

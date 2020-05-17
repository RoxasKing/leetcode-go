package codinginterviews

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	cq := NewCQueue()
	cq.AppendTail(3)
	v := cq.DeleteHead()
	fmt.Println(v)
	v = cq.DeleteHead()
	fmt.Println(v)
}

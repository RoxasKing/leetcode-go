package codinginterviews

import (
	"fmt"
	"testing"
)

func TestCQueue(t *testing.T) {
	cq := Constructor()
	cq.AppendTail(3)
	v := cq.DeleteHead()
	fmt.Println(v)
	v = cq.DeleteHead()
	fmt.Println(v)
}

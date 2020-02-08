package leetcode

import (
	"fmt"
	"testing"
)

func Test_minStack(t *testing.T) {
	minstack := NewMinStack()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	fmt.Println(minstack.GetMin())
	minstack.Pop()
	fmt.Println(minstack.Top())
	fmt.Println(minstack.GetMin())
	minstack.Pop()
	fmt.Println(minstack.Top())
	fmt.Println(minstack.GetMin())
}

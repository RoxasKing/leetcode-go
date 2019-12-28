package My_LeetCode_In_Go

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
}

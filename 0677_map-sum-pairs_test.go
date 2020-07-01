package leetcode

import "testing"

func TestMapSum(t *testing.T) {
	ms := NewMapSum()
	ms.Insert("apple", 3)
	if res := ms.Sum("ap"); res != 3 {
		t.Errorf("ms.Sum(%s)=%d faild,should %d\n", "ap", res, 3)
	}
	ms.Insert("app", 2)
	if res := ms.Sum("ap"); res != 5 {
		t.Errorf("ms.Sum(%s)=%d faild,should %d\n", "ap", res, 5)
	}
}

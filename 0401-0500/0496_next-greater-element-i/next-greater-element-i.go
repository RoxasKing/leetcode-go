package main

// Difficulty:
// Easy

// Tags:
// Hash
// Monotonic Stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := [10001]int{}
	stk := IntStack{}
	for _, num := range nums2 {
		next[num] = -1
		for stk.Len() > 0 && stk.Top() < num {
			next[stk.Pop()] = num
		}
		stk.Push(num)
	}
	out := make([]int, len(nums1))
	for i, num := range nums1 {
		out[i] = next[num]
	}
	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

package main

// Tags:
// Monotone Stack

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	for i := range out {
		out[i] = -1
	}
	s := IntStack{}
	for i, num := range nums {
		for s.Len() > 0 && nums[s.Top()] < num {
			out[s.Pop()] = num
		}
		s.Push(i)
	}
	for _, num := range nums {
		for s.Len() > 0 && nums[s.Top()] < num {
			out[s.Pop()] = num
		}
	}
	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Pop() int {
	last := s.Len() - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}

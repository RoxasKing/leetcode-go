package main

// Tags:
// Monotone Stack
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	out := make([]int, 0, n+1-k)
	stk := IntStack{}
	for i, num := range nums {
		for stk.Len() > 0 && stk.Top() < num {
			stk.Pop()
		}
		stk.Push(num)
		if i > k-1 && stk[0] == nums[i-k] {
			stk = stk[1:]
		}
		if i >= k-1 {
			out = append(out, stk[0])
		}
	}
	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
func (s IntStack) Top() int { return s[s.Len()-1] }

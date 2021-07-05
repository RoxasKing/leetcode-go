package main

// Tags:
// Prefix Sum + Monotone Stack
func maxSumMinProduct(nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)
	for i, num := range nums {
		pSum[i+1] = pSum[i] + num
	}

	stk := IntStack{-1}
	out := 0
	for i, num := range nums {
		for stk.Len() > 1 && num <= nums[stk.Top()] {
			out = Max(out, nums[stk.Pop()]*(pSum[i]-pSum[stk.Top()+1]))
		}
		stk.Push(i)
	}

	for stk.Len() > 1 {
		out = Max(out, nums[stk.Pop()]*(pSum[n]-pSum[stk.Top()+1]))
	}

	return out % (1e9 + 7)
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

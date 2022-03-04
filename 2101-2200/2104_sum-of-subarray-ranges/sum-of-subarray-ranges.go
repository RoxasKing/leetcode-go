package main

// Difficulty:
// Medium

// Tags:
// Monotonic Stack

func subArrayRanges(nums []int) int64 {
	n := len(nums)
	minStk := IntStack{}
	maxStk := IntStack{}
	minLeft := make([]int, n)
	maxLeft := make([]int, n)
	for i, num := range nums {
		for ; minStk.Len() > 0 && nums[minStk.Top()] > num; minStk.Pop() {
		}
		minLeft[i] = -1
		if minStk.Len() > 0 {
			minLeft[i] = minStk.Top()
		}
		minStk.Push(i)
		for ; maxStk.Len() > 0 && nums[maxStk.Top()] <= num; maxStk.Pop() {
		}
		maxLeft[i] = -1
		if maxStk.Len() > 0 {
			maxLeft[i] = maxStk.Top()
		}
		maxStk.Push(i)
	}
	minStk = minStk[:0]
	maxStk = maxStk[:0]
	minRight := make([]int, n)
	maxRight := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		for ; minStk.Len() > 0 && nums[minStk.Top()] >= num; minStk.Pop() {
		}
		minRight[i] = n
		if minStk.Len() > 0 {
			minRight[i] = minStk.Top()
		}
		minStk.Push(i)
		for ; maxStk.Len() > 0 && nums[maxStk.Top()] < num; maxStk.Pop() {
		}
		maxRight[i] = n
		if maxStk.Len() > 0 {
			maxRight[i] = maxStk.Top()
		}
		maxStk.Push(i)
	}
	var sumMax, sumMin int64
	for i, num := range nums {
		sumMax += int64(maxRight[i]-i) * int64(i-maxLeft[i]) * int64(num)
		sumMin += int64(minRight[i]-i) * int64(i-minLeft[i]) * int64(num)
	}
	return sumMax - sumMin
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	i := s.Len() - 1
	out := (*s)[i]
	*s = (*s)[:i]
	return out
}

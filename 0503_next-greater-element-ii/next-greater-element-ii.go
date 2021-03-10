package main

/*
  Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.

  Example 1:
  Input: [1,2,1]
  Output: [2,-1,2]
  Explanation: The first 1's next greater number is 2;
  The number 2 can't find next greater number;
  The second 1's next greater number needs to search circularly, which is also 2.
  Note: The length of given array won't exceed 10000.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/next-greater-element-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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

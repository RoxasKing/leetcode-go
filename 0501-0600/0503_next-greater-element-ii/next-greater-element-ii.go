package main

/*
  Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next greater number for every element in nums.

  The next greater number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, return -1 for this number.

  Example 1:
    Input: nums = [1,2,1]
    Output: [2,-1,2]
    Explanation: The first 1's next greater number is 2;
      The number 2 can't find next greater number.
      The second 1's next greater number needs to search circularly, which is also 2.

  Example 2:
    Input: nums = [1,2,3,4,3]
    Output: [2,3,4,-1,4]

  Constraints:
    1. 1 <= nums.length <= 10^4
    2. -10^9 <= nums[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/next-greater-element-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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

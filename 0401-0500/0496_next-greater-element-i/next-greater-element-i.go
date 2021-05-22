package main

/*
  You are given two integer arrays nums1 and nums2 both of unique elements, where nums1 is a subset of nums2.

  Find all the next greater numbers for nums1's elements in the corresponding places of nums2.

  The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, return -1 for this number.

  Example 1:
    Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
    Output: [-1,3,-1]
    Explanation:
      For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
      For number 1 in the first array, the next greater number for it in the second array is 3.
      For number 2 in the first array, there is no next greater number for it in the second array, so output -1.

  Example 2:
    Input: nums1 = [2,4], nums2 = [1,2,3,4]
    Output: [3,-1]
    Explanation:
      For number 2 in the first array, the next greater number for it in the second array is 3.
      For number 4 in the first array, there is no next greater number for it in the second array, so output -1.

  Constraints:
    1. 1 <= nums1.length <= nums2.length <= 1000
    2. 0 <= nums1[i], nums2[i] <= 10^4
    3. All integers in nums1 and nums2 are unique.
    4. All the integers of nums1 also appear in nums2.

  Follow up: Could you find an O(nums1.length + nums2.length) solution?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/next-greater-element-i
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Monotone Stack

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

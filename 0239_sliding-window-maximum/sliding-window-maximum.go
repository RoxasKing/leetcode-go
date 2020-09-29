package main

/*
  给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
  返回滑动窗口中的最大值。

  进阶：
    你能在线性时间复杂度内解决此题吗？

  提示：
    1 <= nums.length <= 10^5
    -10^4 <= nums[i] <= 10^4
    1 <= k <= nums.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sliding-window-maximum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	out := make([]int, 0, n-k+1)
	stack := MakeStack()
	for i := 0; i < n; i++ {
		for stack.Len() > 0 && stack.Top() < nums[i] {
			_ = stack.Pop()
		}
		stack.Push(nums[i])
		if i > k-1 && nums[i-k] == stack[0] {
			stack = stack[1:]
		}
		if i >= k-1 {
			out = append(out, stack[0])
		}
	}
	return out
}

type Stack []int

func MakeStack() Stack {
	return Stack{}
}

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Top() int {
	return s[s.Len()-1]
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	top := s.Len() - 1
	val := (*s)[top]
	*s = (*s)[:top]
	return val
}

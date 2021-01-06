package main

/*
  给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
  返回滑动窗口中的最大值。

  示例 1：
    输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
    输出：[3,3,5,5,6,7]
    解释：
    滑动窗口的位置                最大值
    ---------------               -----
    [1  3  -1] -3  5  3  6  7       3
     1 [3  -1  -3] 5  3  6  7       3
     1  3 [-1  -3  5] 3  6  7       5
     1  3  -1 [-3  5  3] 6  7       5
     1  3  -1  -3 [5  3  6] 7       6
     1  3  -1  -3  5 [3  6  7]      7

  示例 2：
    输入：nums = [1], k = 1
    输出：[1]

  示例 3：
    输入：nums = [1,-1], k = 1
    输出：[1,-1]

  示例 4：
    输入：nums = [9,11], k = 2
    输出：[11]

  示例 5：
    输入：nums = [4,-2], k = 2
    输出：[4]

  提示：
    1 <= nums.length <= 10^5
    -10^4 <= nums[i] <= 10^4
    1 <= k <= nums.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sliding-window-maximum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	max := make([]int, n+1-k)
	s := IntStack{}
	for i, num := range nums {
		for s.Len() > 0 && s.Top() < num {
			_ = s.Pop()
		}
		s.Push(num)
		if i+1-k > 0 && nums[i-k] == s[0] {
			s = s[1:]
		}
		if i+1-k >= 0 {
			max[i+1-k] = s[0]
		}
	}
	return max
}

type IntStack []int

func (s IntStack) Len() int { return len(s) }

func (s IntStack) Top() int {
	if s.Len() == 0 {
		panic("stack is empty")
	}
	return s[len(s)-1]
}

func (s *IntStack) Pop() int {
	if s.Len() == 0 {
		panic("stack is empty")
	}
	last := len(*s) - 1
	out := (*s)[last]
	*s = (*s)[:last]
	return out
}

func (s *IntStack) Push(x int) { *s = append(*s, x) }

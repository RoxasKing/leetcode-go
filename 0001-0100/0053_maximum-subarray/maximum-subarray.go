package main

/*
  给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

  示例:
    输入: [-2,1,-3,4,-1,2,1,-5,4]
    输出: 6
    解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

  进阶:
    如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxSubArray(nums []int) int {
	sum, max := -1<<31, -1<<31
	for _, num := range nums {
		sum = Max(sum+num, num)
		max = Max(max, sum)
	}
	return max
}

// DFS + Partition
func maxSubArray2(nums []int) int {
	_, _, out, _ := dfs(nums, 0, len(nums)-1)
	return out
}

func dfs(nums []int, l, r int) (lSum, rSum, mSum, aSum int) {
	if l == r {
		return nums[l], nums[l], nums[l], nums[l]
	}
	m := l + (r-l)>>1
	ll, lr, lm, la := dfs(nums, l, m)
	rl, rr, rm, ra := dfs(nums, m+1, r)
	return Max(ll, la+rl), Max(rr, ra+lr), Max(Max(lm, rm), lr+rl), la + ra
}

// Dynamic Programming
func maxSubArrayReturnArray(nums []int) []int {
	sum, max := -1<<31, -1<<31
	var out []int
	var l, r int
	for r = range nums {
		sum += nums[r]
		if sum < nums[r] {
			l, sum = r, nums[r]
		}
		if max < sum {
			max, out = sum, nums[l:r+1]
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

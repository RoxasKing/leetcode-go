package main

/*
  Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

  Example 1:
    Input: nums = [-4,-1,0,3,10]
    Output: [0,1,9,16,100]
    Explanation: After squaring, the array becomes [16,1,0,9,100].
      After sorting, it becomes [0,1,9,16,100].

  Example 2:
    Input: nums = [-7,-3,2,3,11]
    Output: [4,9,9,49,121]

  Constraints:
    1. 1 <= nums.length <= 10^4
    2. -10^4 <= nums[i] <= 10^4
    3. nums is sorted in non-decreasing order.

    Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Merge Sort
func sortedSquares(nums []int) []int {
	n := len(nums)
	m := 0
	for m < n && nums[m] < 0 {
		m++
	}
	reverse(nums[:m])

	for i, num := range nums {
		nums[i] = num * num
	}

	out := make([]int, n)
	i, j, k := 0, m, 0
	for ; i < m && j < n; k++ {
		if nums[i] < nums[j] {
			out[k] = nums[i]
			i++
		} else {
			out[k] = nums[j]
			j++
		}
	}
	copy(out[k:], nums[i:m])
	copy(out[k:], nums[j:n])
	return out
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

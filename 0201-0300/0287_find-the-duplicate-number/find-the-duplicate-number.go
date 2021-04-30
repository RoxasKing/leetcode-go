package main

/*
  Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

  There is only one repeated number in nums, return this repeated number.

  Example 1:
    Input: nums = [1,3,4,2,2]
    Output: 2

  Example 2:
    Input: nums = [3,1,3,4,2]
    Output: 3

  Example 3:
    Input: nums = [1,1]
    Output: 1

  Example 4:
    Input: nums = [1,1,2]
    Output: 1

  Constraints:
    1. 2 <= n <= 10^5
    2. nums.length == n + 1
    3. 1 <= nums[i] <= n
    4. All the integers in nums appear only once except for precisely one integer which appears two or more times.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-the-duplicate-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Two Pointers
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}

// Binary Search
func findDuplicate2(nums []int) int {
	out := 0
	for l, r := 1, len(nums)-1; l <= r; {
		target := (l + r) >> 1
		cnt := 0
		for _, num := range nums {
			if num <= target {
				cnt++
			}
		}
		if cnt <= target {
			l = target + 1
		} else {
			r = target - 1
			out = target
		}
	}
	return out
}

// Bit Manipulation
func findDuplicate3(nums []int) int {
	out := 0
	maxBit := 31
	for (len(nums)-1)>>maxBit == 0 {
		maxBit--
	}
	for bit := 0; bit <= maxBit; bit++ {
		var x, y int
		for i, num := range nums {
			if num&(1<<bit) != 0 {
				x++
			}
			if i > 0 && i&(1<<bit) != 0 {
				y++
			}
		}
		if x > y {
			out |= 1 << bit
		}
	}
	return out
}

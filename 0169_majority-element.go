package leetcode

/*
  给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
  你可以假设数组是非空的，并且给定的数组总是存在多数元素。
  示例 1:
    输入: [3,2,3]
    输出: 3
  示例 2:
    输入: [2,2,1,1,1,2,2]
    输出: 2
*/

func majorityElement(nums []int) int {
	least := len(nums) / 2
	if len(nums)%2 != 0 {
		least++
	}
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] >= least {
			return num
		}
	}
	return 0
}

// Boyer-Moore
func majorityElement2(nums []int) int {
	var count, candidate int
	for _, num := range nums {
		if candidate == num {
			count++
		} else {
			candidate = num
			count = 1
		}
	}
	return candidate
}

// Divide-and-conquer
func majorityElement3(nums []int) int {
	count := func(num, l, r int) int {
		var count int
		for i := l; i <= r; i++ {
			if nums[i] == num {
				count++
			}
		}
		return count
	}
	var recur func(int, int) int
	recur = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		m := l + (r-l)>>1
		left := recur(l, m)
		right := recur(m+1, r)
		if left == right {
			return left
		}
		leftCount := count(left, l, r)
		rightCount := count(right, l, r)
		if leftCount > rightCount {
			return left
		}
		return right
	}
	return recur(0, len(nums)-1)
}

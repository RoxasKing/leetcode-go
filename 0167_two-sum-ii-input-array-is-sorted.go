package main

/*
  给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
  函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

  说明:
    返回的下标值（index1 和 index2）不是从零开始的。
    你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Double Pointer
func twoSumII(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}

// Binary Search
func twoSumII2(numbers []int, target int) []int {
	binarySearch := func(l, r, target int) int {
		for l <= r {
			m := l + (r-l)>>1
			if numbers[m] < target {
				l = m + 1
			} else if numbers[m] > target {
				r = m - 1
			} else {
				return m
			}
		}
		return -1
	}
	for i := 0; i < len(numbers)-1 && numbers[i] <= target/2; i++ {
		if j := binarySearch(i+1, len(numbers)-1, target-numbers[i]); j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

// Hash
func twoSumII3(numbers []int, target int) []int {
	mark := make(map[int]int)
	for i, num := range numbers {
		mark[num] = i
	}
	for i := 0; i < len(numbers) && numbers[i] <= target/2; i++ {
		if j, ok := mark[target-numbers[i]]; ok {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

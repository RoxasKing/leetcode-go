package main

/*
  You are given an integer array nums and you have to return a new counts array. The counts array has the property where counts[i] is the number of smaller elements to the right of nums[i].

  Example 1:
    Input: nums = [5,2,6,1]
    Output: [2,1,1,0]
    Explanation:
      To the right of 5 there are 2 smaller elements (2 and 1).
      To the right of 2 there is only 1 smaller element (1).
      To the right of 6 there is 1 smaller element (1).
      To the right of 1 there is 0 smaller element.

  Example 2:
    Input: nums = [-1]
    Output: [0]

  Example 3:
    Input: nums = [-1,-1]
    Output: [0,0]

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. -10^4 <= nums[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Merge Sort
func countSmaller(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}

	merge(nums, out, idxs, 0, n-1)

	return out
}

func merge(nums, out, idxs []int, l, r int) {
	if l == r {
		return
	}

	m := (l + r) >> 1
	merge(nums, out, idxs, l, m)
	merge(nums, out, idxs, m+1, r)

	if nums[idxs[m]] <= nums[idxs[m+1]] {
		return
	}

	tmp := make([]int, r+1-l)

	i, j, k := l, m+1, 0
	for ; i <= m && j <= r; k++ {
		if nums[idxs[i]] <= nums[idxs[j]] {
			out[idxs[i]] += j - (m + 1)
			tmp[k] = idxs[i]
			i++
		} else {
			tmp[k] = idxs[j]
			j++
		}
	}

	for ; i <= m; i, k = i+1, k+1 {
		out[idxs[i]] += r - m
		tmp[k] = idxs[i]
	}
	copy(tmp[k:k+r+1-j], idxs[j:r+1])
	copy(idxs[l:r+1], tmp)
}

package main

// Tags:
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

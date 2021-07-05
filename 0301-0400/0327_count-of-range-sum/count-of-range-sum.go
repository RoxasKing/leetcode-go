package main

// Tags:
// Hash
func countRangeSum(nums []int, lower int, upper int) int {
	dict := make(map[int]int)
	n := len(nums)
	count, sum := 0, 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if lower <= sum && sum <= upper {
			count++
		}
		for j := lower; j <= upper; j++ {
			if v, ok := dict[sum-j]; ok {
				count += v
			}
		}
		dict[sum]++
	}
	return count
}

// Merge Sort
func countRangeSum2(nums []int, lower int, upper int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := range nums {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	prefixTmp := make([]int, n+1)
	return mergeCount(prefixSum, prefixTmp, lower, upper, 0, n)
}

func mergeCount(sum, tmp []int, lower, upper, l, r int) int {
	if l == r {
		return 0
	}

	m := l + (r-l)>>1
	count := mergeCount(sum, tmp, lower, upper, l, m) + mergeCount(sum, tmp, lower, upper, m+1, r)
	copy(tmp[l:r+1], sum[l:r+1])

	// count pair
	i, j := m+1, m+1
	for _, num := range tmp[l : m+1] {
		for i <= r && tmp[i]-num < lower {
			i++
		}
		for j <= r && tmp[j]-num <= upper {
			j++
		}
		count += j - i
	}

	// merge tmp[l:m+1] and tmp[m+1:r] to arr[l:r+1]
	p0, p1, p2 := l, l, m+1
	for p1 <= m && p2 <= r {
		if tmp[p1] < tmp[p2] {
			sum[p0] = tmp[p1]
			p1++
		} else {
			sum[p0] = tmp[p2]
			p2++
		}
		p0++
	}
	copy(sum[p0:r+1], tmp[p1:m+1])
	copy(sum[p0:r+1], tmp[p2:r+1])

	return count
}

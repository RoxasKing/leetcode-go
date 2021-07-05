package main

// Tags:
// Sliding Window + Greedy Algorithm + Two Pointers
func minOperations(nums []int, x int) int {
	n := len(nums)
	l, r := 0, n-1
	sum := 0
	out := 1<<31 - 1

	for ; l < n; l++ {
		sum += nums[l]
		if sum == x {
			out = Min(out, l+1)
			break
		}

		if sum > x {
			sum -= nums[l]
			l--
			break
		}
	}

	if l == n && sum < x {
		return -1
	}

	for {
		for l < r && sum < x {
			sum += nums[r]
			r--
		}

		if sum == x {
			out = Min(out, l+1+n-r-1) // l+1: leftmost number, n-r-1: rightmost number
		}

		if l < 0 {
			break
		}

		sum -= nums[l]
		l--
	}

	if out == 1<<31-1 {
		return -1
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Sliding Window + Hash
func minOperations2(nums []int, x int) int {
	out := 1<<31 - 1
	n := len(nums)
	sum := 0
	dict := make(map[int]int)
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum == x {
			out = Min(out, i+1)
		}
		if sum >= x {
			break
		}
		dict[sum] = i
	}
	sum = 0
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		if sum == x {
			out = Min(out, n-i)
		}
		if idx, ok := dict[x-sum]; ok && idx < i {
			out = Min(out, idx+1+n-i)
		}
		if sum >= x {
			break
		}
	}
	if out == 1<<31-1 {
		return -1
	}
	return out
}

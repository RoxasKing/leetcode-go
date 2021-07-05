package main

import "sort"

func nextGreaterElement(n int) int {
	arr := []int{}
	for i := n; i > 0; i /= 10 {
		arr = append(arr, i%10)
	}

	reverse(arr)

	m := len(arr)
	i := m - 1
	for i-1 >= 0 && arr[i-1] >= arr[i] {
		i--
	}

	reverse(arr[i:])

	if i == 0 {
		return -1
	}

	j := sort.Search(m-i, func(j int) bool { return arr[j+i] > arr[i-1] }) + i
	arr[i-1], arr[j] = arr[j], arr[i-1]

	out := 0
	for _, num := range arr {
		if (1<<31-1-num)/10 < out {
			return -1
		}
		out = out*10 + num
	}
	return out
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

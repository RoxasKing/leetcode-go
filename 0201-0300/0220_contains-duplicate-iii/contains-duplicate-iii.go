package main

import "sort"

// Bucket Sort
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	buckets := make(map[int]int)
	for i, a := range nums {
		bucket := getID(a, t+1)
		if _, ok := buckets[bucket]; ok {
			return true
		}
		if b, ok := buckets[bucket-1]; ok && Abs(a-b) <= t {
			return true
		}
		if b, ok := buckets[bucket+1]; ok && Abs(a-b) <= t {
			return true
		}
		buckets[bucket] = a
		if i >= k {
			delete(buckets, getID(nums[i-k], t+1))
		}
	}
	return false
}

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

// Binary Search
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	n := len(nums)
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return nums[idxs[i]] < nums[idxs[j]]
	})
	for i := range nums {
		l := sort.Search(n, func(j int) bool {
			return nums[idxs[j]] >= nums[i]-t
		})
		r := sort.Search(n, func(j int) bool {
			return nums[idxs[j]] > nums[i]+t
		}) - 1
		for j := l; j <= r; j++ {
			if idxs[j] == i {
				continue
			}
			if Abs(i-idxs[j]) <= k {
				return true
			}
		}
	}
	return false
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

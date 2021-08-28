package main

// Tags:
// Greedy
// Monotone Stack

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	out := []int{}
	n1, n2 := len(nums1), len(nums2)
	for i := Max(0, k-n2); i <= Min(n1, k); i++ {
		if tmp := merge(chooseMax(nums1, i), chooseMax(nums2, k-i)); isBigger(tmp, out) {
			out = tmp
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func chooseMax(nums []int, choose int) []int {
	del := len(nums) - choose
	out := []int{}
	for _, num := range nums {
		for len(out) > 0 && out[len(out)-1] < num && del > 0 {
			out = out[:len(out)-1]
			del--
		}
		out = append(out, num)
	}
	return out[:choose]
}

func merge(nums1, nums2 []int) []int {
	out := make([]int, 0, len(nums1)+len(nums2))
	for len(nums1) > 0 && len(nums2) > 0 {
		if isBigger(nums1, nums2) {
			out = append(out, nums1[0])
			nums1 = nums1[1:]
		} else {
			out = append(out, nums2[0])
			nums2 = nums2[1:]
		}
	}
	out = append(out, nums1...)
	out = append(out, nums2...)
	return out
}

func isBigger(nums1, nums2 []int) bool {
	n1, n2 := len(nums1), len(nums2)
	for i := 0; i < Min(n1, n2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] > nums2[i]
		}
	}
	return n1 > n2
}

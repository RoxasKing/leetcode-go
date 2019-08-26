package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	num1 := nums1[:m]
	num2 := nums2[:n]
	out := make([]int, 0)
	i, j := 0, 0
	for i < len(num1) || j < len(num2) {
		if i < len(num1) && j < len(num2) {
			if num1[i] < num2[j] {
				out = append(out, num1[i])
				i++
			} else {
				out = append(out, num2[j])
				j++
			}
		} else if i < len(num1) && j >= len(num2) {
			out = append(out, num1[i])
			i++
		} else {
			out = append(out, num2[j])
			j++
		}
	}
	for i := range out {
		nums1[i] = out[i]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m, n := 3, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}

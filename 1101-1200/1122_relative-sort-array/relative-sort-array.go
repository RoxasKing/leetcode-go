package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
	count := make([]int, 1001)
	for _, num := range arr1 {
		count[num]++
	}
	out := make([]int, 0, len(arr1))
	for _, num := range arr2 {
		for count[num] > 0 {
			out = append(out, num)
			count[num]--
		}
	}
	for i := range count {
		for count[i] > 0 {
			out = append(out, i)
			count[i]--
		}
	}
	return out
}

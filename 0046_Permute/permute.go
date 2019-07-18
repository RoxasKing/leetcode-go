package main

import "fmt"

func permute(nums []int) [][]int {
	size := len(nums)
	if size < 1 {
		return [][]int{nums}
	}
	out := [][]int{nums[:1]}
	tmps := [][]int{}
	tmp := []int{}

	for i := 1; i < size; i++ {
		for _, elem := range out {
			for j := 0; j <= len(elem); j++ {
				if j == 0 {
					tmp = append(tmp, nums[i])
					tmp = append(tmp, elem...)
				} else if j == len(elem) {
					tmp = append(tmp, elem...)
					tmp = append(tmp, nums[i])
				} else {
					tmp = append(tmp, elem[:j]...)
					tmp = append(tmp, nums[i])
					tmp = append(tmp, elem[j:]...)
				}
				tmps = append(tmps, tmp)
				tmp = []int{}
			}
		}
		out = tmps
		tmps = nil
	}
	return out
}

func main() {
	in := []int{1, 2, 3}
	fmt.Println(permute(in))
}

package main

import "fmt"

func permute(nums []int) [][]int {
	size := len(nums)
	if size < 1 {
		return [][]int{nums}
	}
	out := [][]int{}
	tmps := [][]int{}
	tmp := []int{}

	for _, num := range nums {
		for _, elem := range out {
			for i := 0; i < len(elem); i++ {
				if i == 0 {
					tmp = append([]int{num}, elem...)
				} else if i == len(elem)-1 {
					tmp = append(elem, nums[i])
				} else {
					tmp = append(elem[:i], nums[i])
					tmp = append(tmp, nums[i+1])
				}
				tmps = append(tmps, tmp)
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

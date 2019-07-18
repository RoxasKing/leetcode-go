package main

import (
	"fmt"
	"strconv"
)

func permuteUnique(nums []int) [][]int {
	size := len(nums)
	if size < 1 {
		return [][]int{nums}
	}
	out := [][]int{nums[:1]}
	tmps := [][]int{}
	tmp := []int{}
	dict := make(map[string]bool)

	for i := 1; i < size; i++ {
		for _, elem := range out {
			for j := 0; j <= len(elem); j++ {
				if j == 0 {
					tmp = append(tmp, nums[i])
					tmp = append(tmp, elem...)
				} else if j == len(elem) {
					if nums[i] == elem[j-1] {
						break
					}
					tmp = append(tmp, elem...)
					tmp = append(tmp, nums[i])
				} else {
					if nums[i] == elem[j-1] {
						continue
					}
					tmp = append(tmp, elem[:j]...)
					tmp = append(tmp, nums[i])
					tmp = append(tmp, elem[j:]...)
				}
				if i != size-1 {
					tmps = append(tmps, tmp)
				} else {
					str := ""
					for _, elem := range tmp {
						str += strconv.Itoa(elem)
					}
					if _, ok := dict[str]; !ok {
						tmps = append(tmps, tmp)
						dict[str] = true
					}
				}
				tmp = []int{}
			}
		}
		out = tmps
		tmps = nil
	}
	return out
}

func main() {
	in := []int{1, 1, 2, 2}
	fmt.Println(permuteUnique(in))
}

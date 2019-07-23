package main

import (
	"fmt"
	"sort"
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

func permuteUnique2(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	// vector 是一组可能的解答
	vector := make([]int, n)
	// taken[i] == true 表示 nums[i] 已经出现在了 vector 中
	taken := make([]bool, n)

	var ans [][]int

	makePermutation(0, n, nums, vector, taken, &ans)

	return ans
}

// 这个方式和我原来的方式相比，
// 增加了比较的次数
// 但是，减少了切片生成的次数。
// 所以，速度会快一些。
func makePermutation(cur, n int, nums, vector []int, taken []bool, ans *[][]int) {
	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	used := make(map[int]bool, n-cur)

	for i := 0; i < n; i++ {

		if !taken[i] && !used[nums[i]] {
			used[nums[i]] = true

			// 准备使用 nums[i]，所以，taken[i] == true
			taken[i] = true
			// NOTICE: 是 vector[cur]
			vector[cur] = nums[i]

			makePermutation(cur+1, n, nums, vector, taken, ans)

			// 下一个循环中
			// vector[cur] = nums[i+1]
			// 所以，在这个循环中，恢复 nums[i] 自由
			taken[i] = false
		}
	}
}

func main() {
	in := []int{1, 1, 2, 2}
	fmt.Println(permuteUnique(in))
	fmt.Println(permuteUnique2(in))
}

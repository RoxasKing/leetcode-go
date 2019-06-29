package main

import "fmt"

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return nil
	}
	var out [][]int
	for i := 0; i < size; i++ {
		for j := 0; j < size-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	i := 0
	j := 1
	k := 2
	l := 3
	for l < size-1 {
		sum := nums[i] + nums[j] + nums[k] + nums[l]
		if sum < target {
			if i < j && j < k && k < l && l < size-1 && nums[size-4]+nums[size-3]+nums[size-2]+nums[size-1] < target {
				break
			}
			if i < j && j < k && k < l && l < size-1 && nums[i]+nums[size-3]+nums[size-2]+nums[size-1] < target {
				i++
				j = i + 1
				k = j + 1
				l = k + 1
				for l < size-1 && nums[l] == nums[l-1] {
					i++
					j++
					k++
					l++
				}
			}
			if i < j && j < k && k < l && l < size-1 && nums[i]+nums[j]+nums[size-2]+nums[size-1] < target {
				j++
				k = j + 1
				l = k + 1
				for l < size-1 && nums[l] == nums[l-1] {
					j++
					k++
					l++
				}
			}
			if i < j && j < k && k < l && l < size-1 && nums[i]+nums[j]+nums[k]+nums[size-1] < target {
				k++
				l = k + 1
				for l < size-1 && nums[l] == nums[l-1] {
					k++
					l++
				}
			}
			l++
			for l < size-1 && nums[l] == nums[l-1] {
				l++
				for l < size-1 && nums[l] == nums[l-1] {
					l++
				}
			}
		} else if sum > target {
			k++
			l = k + 1
			if l < size-1 {
				continue
			} else {
				j++
				k = j + 1
				l = k + 1
				if l < size-1 {
					continue
				} else {
					i++
					j = i + 1
					k = j + 1
					l = k + 1
					if l < size-1 {
						continue
					} else {
						break
					}
				}
			}
		} else {
			out = append(out, []int{nums[i], nums[j], nums[k], nums[l]})
			l++
			for l < size-1 && nums[l] == nums[l-1] {
				l++
			}
			continue
		}
	}

	return out
}

func main() {
	nums := []int{0, 4, -5, 2, -2, 4, 2, -1, 4}
	target := 12
	fmt.Println(fourSum(nums, target))
}

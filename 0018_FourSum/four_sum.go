package main

import "fmt"

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	if size < 4 {
		return nil
	}
	var out [][]int
	QuickSort(nums)
	for i := 0; i < size-3; i++ {
		// 如果当前下标 i 的值与 i-1 相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 如果当前下标 i 最小和仍大于目标值
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 如果当前下标 i 最大和仍小于目标值
		if nums[i]+nums[size-1]+nums[size-2]+nums[size-3] < target {
			continue
		}
		//
		for j := i + 1; j < size-2; j++ {
			if j-i > 1 && nums[j] == nums[j-1] {
				continue
			}
			// 如果当前下标 j 最小和仍大于目标值
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// 如果当前下标 i 最大和仍小于目标值
			if nums[i]+nums[j]+nums[size-1]+nums[size-2] < target {
				continue
			}
			k, l := j+1, size-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					out = append(out, []int{nums[i], nums[j], nums[k], nums[l]})
					l--
					for l > k && nums[l] == nums[l+1] {
						l--
					}
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
				} else if sum > target {
					l--
					for l > k && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
				}
			}
		}
	}
	return out
}

func main() {
	nums := []int{0, 0, 0, 0}
	//nums := []int{0, 4, -5, 2, -2, 4, 2, -1, 4}
	//target := 12
	//nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	//nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	//target := 0
	//nums := []int{-1, -5, -5, -3, 2, 5, 0, 4}
	//target := -7
	fmt.Println(fourSum(nums, target))
}

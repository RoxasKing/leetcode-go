package main

import "fmt"

func searchRange(nums []int, target int) []int {
	size := len(nums)
	out := []int{-1, -1}

	i := 0
	for ; i < size; i++ {
		if nums[i] == target {
			out[0] = i
			break
		}
	}
	if i < size {
		for i < size-1 && nums[i+1] == target {
			i++
		}
		out[1] = i
	}
	return out
}
func searchRange2(nums []int, target int) []int {
	out := []int{-1, -1}
	size := len(nums)
	if size < 1 {
		return out
	}
	head := searchBonde(nums, target, false) // 查找左边界
	tail := searchBonde(nums, target, true)  // 查找右边界
	if nums[head] == target {
		out[0] = head
	} else if head+1 < size && nums[head+1] == target {
		out[0] = head + 1
	}
	if nums[tail] == target {
		out[1] = tail
	} else if tail-1 >= 0 && nums[tail-1] == target {
		out[1] = tail - 1
	}
	return out
}
func searchBonde(nums []int, target int, flag bool) int {
	head, tail := 0, len(nums)-1
	for head < tail {
		mid := (head + tail) / 2
		if target == nums[mid] {
			if flag {
				// 查找右边界，右移
				head = mid + 1
			} else {
				// 查找左边界，左移
				tail = mid - 1
			}
		} else if target > nums[mid] {
			head = mid + 1
		} else if target < nums[mid] {
			tail = mid - 1
		}
	}
	return head
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	//nums := []int{1, 1, 2, 3, 4, 5, 5}
	//target := 7
	//nums := []int{1}
	//target := 1
	fmt.Println(searchRange2(nums, target))
}

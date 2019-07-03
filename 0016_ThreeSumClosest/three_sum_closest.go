package main

import "fmt"

func threeSumClosest(nums []int, target int) int {
	size := len(nums)
	out := nums[size-1] + nums[size-2] + nums[size-3]
	QuickSort(nums)

	i := 0
	for i < size-1 {
		j := i + 1
		k := size - 1
		for j < k {
			closest := out - target
			sum := nums[i] + nums[j] + nums[k]
			deviation := sum - target
			if deviation < 0 {
				if closest < 0 && deviation > closest || closest >= 0 && -deviation < closest {
					out = sum
				}
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if deviation > 0 {
				if closest < 0 && -deviation > closest || closest >= 0 && deviation < closest {
					out = sum
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else {
				return sum
			}
		}
		i++
		for i < size-1 && nums[i] == nums[i-1] {
			i++
		}
	}
	return out
}

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

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	//nums := []int{-1, 2, 1, -4}
	fmt.Println(nums)
	fmt.Println(threeSumClosest(nums, 1))
}

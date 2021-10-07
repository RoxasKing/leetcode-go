package main

// Difficulty:
// Medium

// Tags:
// Sort

func findDuplicates(nums []int) []int {
	out := []int{}
	for _, num := range nums {
		i := Abs(num)
		if nums[i-1] < 0 {
			out = append(out, i)
		} else {
			nums[i-1] = -nums[i-1]
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

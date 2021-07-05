package main

// Tags:
// Dynamic Programming

func maxAlternatingSum(nums []int) int64 {
	add, del := int64(-1<<31), int64(0)
	for _, num := range nums {
		add = Max(add, del+int64(num))
		del = Max(del, add-int64(num))
	}
	return add
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

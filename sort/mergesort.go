package sort

func MergeSort(array []int) {
	gap := 2
	for ; gap <= len(array); gap *= 2 {
		var i int
		for ; i+gap < len(array); i = i + gap {
			copy(array[i:i+gap], merge(array[i:i+gap/2], array[i+gap/2:i+gap]))
		}
		copy(array[i:], merge(array[i:i+gap/2], array[i+gap/2:]))
	}
	copy(array, merge(array[:gap/2], array[gap/2:]))
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}

func MergeSortInRecur(array []int) {
	var divide func([]int) []int
	divide = func(array []int) []int {
		if len(array) < 2 {
			return array
		}
		mid := len(array) / 2
		return merge(divide(array[0:mid]), divide(array[mid:]))
	}
	copy(array, divide(array))
}

package sort

func MergeSort(array []int) {
	if len(array) < 2 {
		return
	}
	gap := 2
	var i int
	for {
		for i = 0; i+gap < len(array); i = i + gap {
			copy(array[i:i+gap], merge(array[i:i+gap/2], array[i+gap/2:i+gap]))
		}
		copy(array[i:], merge(array[i:i+gap/2], array[i+gap/2:]))
		gap *= 2
		if gap > len(array) {
			copy(array, merge(array[:gap/2], array[gap/2:]))
			break
		}
	}
}

func mergeSort_recur(array []int) {
	copy(array, divide(array))
}

func divide(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2
	return merge(divide(array[0:mid]), divide(array[mid:]))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

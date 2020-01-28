package sort

func CountingSort(array []int) {
	max := array[0]
	for i := range array {
		if array[i] > max {
			max = array[i]
		}
	}
	c := make([]int, max+1)
	for i := range array {
		c[array[i]]++
	}
	for i := 1; i <= max; i++ {
		c[i] = c[i] + c[i-1]
	}
	newArray := make([]int, len(array))
	for i := len(array) - 1; i >= 0; i-- {
		newArray[c[array[i]]-1] = array[i]
		c[array[i]]--
	}
	copy(array, newArray)
}

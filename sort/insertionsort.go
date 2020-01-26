package sort

func InsertionSort(array []int) {
	for i := range array {
		pre, cur := i-1, array[i]
		for ; pre >= 0 && array[pre] > cur; pre-- {
			array[pre+1] = array[pre]
		}
		array[pre+1] = cur
	}
}

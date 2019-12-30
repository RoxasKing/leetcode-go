package sort

func InsertionSort(array []int) {
	var pre, tmp int
	for i := range array {
		pre, tmp = i-1, array[i]
		for pre >= 0 && array[pre] > tmp {
			array[pre+1] = array[pre]
			pre -= 1
		}
		array[pre+1] = tmp
	}
}

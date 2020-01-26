package sort

func HeapSort(array []int) {
	nodeRise := func(lastIndex int) {
		for i := len(array)/2 - 1; i >= 0; i-- {
			son := i*2 + 1
			if son > lastIndex {
				continue
			}
			if son < lastIndex && array[son] < array[son+1] {
				son++
			}
			if array[son] > array[i] {
				array[son], array[i] = array[i], array[son]
			}
		}
	}
	for i := range array {
		nodeRise(len(array) - 1 - i)
		array[0], array[len(array)-1-i] = array[len(array)-1-i], array[0]
	}
}

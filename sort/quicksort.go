package sort

func QuickSort(array []int) {
	if len(array) < 2 {
		return
	}
	midElem, mid, tail := array[0], 0, len(array)-1
	for i := 1; i <= tail; {
		if array[i] > midElem {
			array[i], array[tail], tail = array[tail], array[i], tail-1
		} else {
			array[i], array[mid], mid, i = array[mid], array[i], mid+1, i+1
		}
	}
	QuickSort(array[:mid])
	QuickSort(array[mid+1:])
}

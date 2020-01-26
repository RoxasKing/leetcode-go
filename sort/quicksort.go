package sort

func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	mid, head, tail := array[0], 0, len(array)-1
	for i := 1; i <= tail; {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	QuickSort(array[:head])
	QuickSort(array[head+1:])
}

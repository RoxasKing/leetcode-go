package Algorithm

import (
	"fmt"
	"testing"
)

func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}

func TestQuickSort(t *testing.T) {
	data := []int{5, 6, 7, 3, 8, 1, 9, 0, 3, 5}
	fmt.Println(data)
	QuickSort(data)
	fmt.Println(data)
}

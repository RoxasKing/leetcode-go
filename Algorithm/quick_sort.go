package main

import "fmt"

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

func main() {
	list := []int{8, 56, 4, 3, 5, 78, 12, 5, 0, 6, 3, 12, 12, 56}
	QuickSort(list)
	fmt.Println(list)
}

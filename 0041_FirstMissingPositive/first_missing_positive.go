package main

import "fmt"

func firstMissingPositive(nums []int) int {
	size := len(nums)
	dict := make(map[int]bool)
	for i := 0; i < size; i++ {
		dict[nums[i]] = true
	}
	out := 1
	for {
		if _, ok := dict[out]; !ok {
			return out
		}
		out++
	}
}

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
	//in := []int{1, 2, 0}
	//in := []int{3, 4, -1, 1}
	//in := []int{7, 8, 9, 11, 12}
	//in := []int{-1}
	in := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(in))
}

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	QuickSort(candidates)
	out := combination(candidates, target)
	return out
}

func combination(candidates []int, target int) [][]int {
	var out [][]int
	size := len(candidates)
	for i := 0; i < size; i++ {
		if candidates[i] == target {
			out = append(out, []int{candidates[i]})
		} else if target-candidates[i] >= candidates[i] {
			temp := combination(candidates[i:], target-candidates[i])
			if temp != nil {
				for _, elem := range temp {
					out = append(out, append(elem, candidates[i]))
				}
			}
		}
	}
	return out
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
	candidates := []int{8, 7, 4, 3}
	target := 11
	fmt.Println(combinationSum(candidates, target))
}

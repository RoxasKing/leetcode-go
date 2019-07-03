package main

import "fmt"

func combinationSum2(candidates []int, target int) [][]int {
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
			temp := combination(candidates[i+1:], target-candidates[i])
			if temp != nil {
				for _, elem := range temp {
					out = append(out, append(elem, candidates[i]))
				}
			}
		}
		for i+1 < size && candidates[i+1] == candidates[i] {
			i++
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
	//candicates := []int{10, 1, 2, 7, 6, 1, 5}
	//candicates := []int{2, 5, 2, 1, 2}
	//target := 5
	//candicates := []int{2, 2, 2}
	//target := 2
	candicates := []int{4, 4, 2, 1, 4, 2, 2, 1, 3}
	target := 6
	fmt.Println(combinationSum2(candicates, target))
}

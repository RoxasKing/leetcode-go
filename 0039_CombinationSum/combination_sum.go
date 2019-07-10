package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	// 快排数组，从大到小顺序
	QuickSort(candidates)
	// 找出无重复数组
	out := combination(candidates, target)
	return out
}

func combination(candidates []int, target int) [][]int {
	// 定义一个二维数组用于输出
	var out [][]int
	size := len(candidates)
	for i := 0; i < size; i++ {
		if candidates[i] == target {
			// 如果当前数组元素直接满足目标值，则作为单元素数组add到二维数组中
			out = append(out, []int{candidates[i]})
		} else if target-candidates[i] >= candidates[i] {
			// 如果当前数组元素小于等于目标值,则从当前下标开始作为一个 新切片，和 目标值-当前下标元素值 作为新目标值，
			// 寻找符合要求的二维数组
			temp := combination(candidates[i:], target-candidates[i])
			// 如果结果集不为空
			if temp != nil {
				for _, elem := range temp {
					// 将结果合并为新的数组加入到二维数组中
					out = append(out, append([]int{candidates[i]}, elem...))
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

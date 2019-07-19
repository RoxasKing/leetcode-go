package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	out := make([][]string, 0, 1)
	dict := make(map[string]int)
	point := 0
	for i := 0; i < len(strs); i++ {
		bytes := []byte(strs[i])
		QuickSort(bytes)
		fmt.Println(string(bytes))
		if value, ok := dict[string(bytes)]; !ok {
			out[point] = make([]string, 0, 1)
			out[point] = append(out[point], strs[i])
			dict[string(bytes)] = point
			point++
		} else {
			out[value] = append(out[value], strs[i])
		}
	}
	return out
}

func QuickSort(data []byte) {
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
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(input))
}

package main

import "fmt"

func groupAnagrams2(strs []string) [][]string {
	out := [][]string{}
	// 用于记录字符排序后相同的字符串数组
	dicts := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		QuickSort(bytes)
		key := string(bytes)
		dict := dicts[key]
		dict = append(dict, str)
		dicts[key] = dict
	}
	for _, dict := range dicts {
		out = append(out, dict)
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
	fmt.Println(groupAnagrams2(input))
}

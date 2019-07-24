package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	out := [][]string{}
	// 用于记录字符排序后的字符串，并记忆在输出多维数组的下标位置
	dict := make(map[string]int)
	for point, i := 0, 0; i < len(strs); i++ {
		bytes := []byte(strs[i])
		QuickSort(bytes)
		str := string(bytes)
		if value, ok := dict[str]; !ok {
			// 如果在字典中不存在，则新建一个 string 数组，append到 out 中，并在字典中做记录
			out = append(out, []string{strs[i]})
			dict[str] = point
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

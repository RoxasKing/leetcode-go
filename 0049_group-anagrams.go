package leetcode

import "sort"

/*
  给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
*/

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	sortStr := func(str string) string {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		return string(bytes)
	}
	for _, str := range strs {
		key := sortStr(str)
		dict[key] = append(dict[key], str)
	}
	out := make([][]string, 0, len(dict))
	for _, strs := range dict {
		out = append(out, strs)
	}
	return out
}

func groupAnagrams2(strs []string) [][]string {
	var out [][]string
	mark := make(map[[26]byte]int)
	for _, str := range strs {
		var key [26]byte
		for _, b := range str {
			key[b-'a']++
		}
		if j, ok := mark[key]; ok {
			out[j] = append(out[j], str)
		} else {
			out = append(out, []string{str})
			mark[key] = len(mark)
		}
	}
	return out
}

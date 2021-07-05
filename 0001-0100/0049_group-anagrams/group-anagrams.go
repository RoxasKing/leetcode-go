package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)
	for _, str := range strs {
		key := sortedString(str)
		dict[key] = append(dict[key], str)
	}
	out := make([][]string, 0, len(dict))
	for _, v := range dict {
		out = append(out, v)
	}
	return out
}

func sortedString(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	return string(bs)
}

func groupAnagrams2(strs []string) [][]string {
	var out [][]string
	mark := make(map[[26]byte]int)
	for _, str := range strs {
		var key [26]byte
		for _, b := range str {
			key[b-'a']++
		}
		if i, ok := mark[key]; ok {
			out[i] = append(out[i], str)
		} else {
			out = append(out, []string{str})
			mark[key] = len(mark)
		}
	}
	return out
}

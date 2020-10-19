package main

import "sort"

/*
  给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

  示例:
    输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
    输出:
    [
      ["ate","eat","tea"],
      ["nat","tan"],
      ["bat"]
    ]

  说明：
    所有输入均为小写字母。
    不考虑答案输出的顺序。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/group-anagrams
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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

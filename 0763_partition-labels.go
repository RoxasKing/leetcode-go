package main

/*
  字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一个字母只会出现在其中的一个片段。返回一个表示每个字符串片段的长度的列表。

  提示：
    S的长度在[1, 500]之间。
    S只包含小写字母 'a' 到 'z' 。
*/

// Greedy Algorithm
func partitionLabels(S string) []int {
	var last [26]int
	for i := range S {
		last[S[i]-'a'] = i
	}
	var out []int
	var l, r int
	for i := range S {
		r = Max(r, last[S[i]-'a'])
		if i == r {
			out = append(out, r+1-l)
			l = r + 1
		}
	}
	return out
}

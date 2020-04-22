package leetcode

/*
  给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
  换句话说，第一个字符串的排列之一是第二个字符串的子串。
*/

func checkInclusion(s1 string, s2 string) bool {
	have, need := [26]int{}, [26]int{}
	for _, c := range s1 {
		need[c-97]++
	}
	var l, r, count int
	for r < len(s2) {
		if need[s2[r]-97] == 0 {
			have = [26]int{}
			l, r, count = r+1, r+1, 0
			continue
		}
		count++
		have[s2[r]-97]++
		for l <= r && have[s2[r]-97] > need[s2[r]-97] {
			count--
			have[s2[l]-97]--
			l++
		}
		if count == len(s1) {
			return true
		}
		r++
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	have, need := [26]int{}, [26]int{}
	for _, c := range s1 {
		need[c-97]++
	}
	var l, r int
	for r < len(s2) {
		have[s2[r]-97]++
		for l <= r && have[s2[r]-97] > need[s2[r]-97] {
			have[s2[l]-97]--
			l++
		}
		if r+1-l == len(s1) {
			return true
		}
		r++
	}
	return false
}

func checkInclusion3(s1 string, s2 string) bool {
	save := [26]int{}
	for _, c := range s1 {
		save[c-97]++
	}
	var l, r int
	for r < len(s2) {
		save[s2[r]-97]--
		for l <= r && save[s2[r]-97] < 0 {
			save[s2[l]-97]++
			l++
		}
		if r+1-l == len(s1) {
			return true
		}
		r++
	}
	return false
}

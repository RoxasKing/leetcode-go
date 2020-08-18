package main

/*
  给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

  说明:
    你可以假设字符串只包含小写字母。
  进阶:
    如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

func isAnagram(s string, t string) bool {
	sc, tc := [26]int{}, [26]int{}
	for i := range s {
		sc[int(s[i]-'a')]++
	}
	for i := range t {
		tc[int(t[i]-'a')]++
	}
	for i := range s {
		if sc[int(s[i]-'a')] != tc[int(s[i]-'a')] {
			return false
		}
	}
	for i := range t {
		if sc[int(t[i]-'a')] != tc[int(t[i]-'a')] {
			return false
		}
	}
	return true
}

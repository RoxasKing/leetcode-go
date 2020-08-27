package main

/*
  给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
  给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
  https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png
*/

// Iteration
func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	dict := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	out := dict[int(digits[0]-'0')]
	for i := 1; i < len(digits); i++ {
		letters := dict[int(digits[i]-'0')]
		tmp := make([]string, 0, len(out)*len(letters))
		for _, str := range out {
			for _, letter := range letters {
				tmp = append(tmp, str+letter)
			}
		}
		out = tmp
	}
	return out
}

// Backtracking
func letterCombinations2(digits string) []string {
	if digits == "" {
		return nil
	}
	dict := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	var out []string
	var backtrack func(int, string)
	backtrack = func(index int, str string) {
		if len(str) == len(digits) {
			out = append(out, str)
			return
		}
		letters := dict[int(digits[index]-'0')]
		for _, letter := range letters {
			backtrack(index+1, str+letter)
		}
	}
	backtrack(0, "")
	return out
}
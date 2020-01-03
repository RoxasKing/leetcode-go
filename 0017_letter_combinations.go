package My_LeetCode_In_Go

/*
  给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
  给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
  https://assets.leetcode-cn.com/aliyun-lc-upload/original_images/17_telephone_keypad.png
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	dict := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	out := dict[digits[0]]

	for i := 1; i < len(digits); i++ {
		var tmp []string
		for j := range out {
			for k := range dict[digits[i]] {
				tmp = append(tmp, out[j]+dict[digits[i]][k])
			}
		}
		out = tmp
	}

	return out
}

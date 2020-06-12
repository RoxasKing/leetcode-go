package codinginterviews

import "strings"

/*
  输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

  说明：
    无空格字符构成一个单词。
    输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
    如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseWords(s string) string {
	var strs []string
	var l int
	for l < len(s) {
		for l < len(s) && s[l] == ' ' {
			l++
		}
		r := l
		for r < len(s) && s[r] != ' ' {
			r++
		}
		if l < r {
			strs = append(strs, s[l:r])
		}
		l = r
	}
	for i := 0; i < len(strs)>>1; i++ {
		strs[i], strs[len(strs)-1-i] = strs[len(strs)-1-i], strs[i]
	}
	return strings.Join(strs, " ")
}

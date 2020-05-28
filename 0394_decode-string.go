package leetcode

import (
	"strconv"
	"strings"
)

/*
  给定一个经过编码的字符串，返回它解码后的字符串。
  编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
  你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
  此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/decode-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func decodeString(s string) string {
	var stack []string
	var index int
	for index < len(s) {
		if '0' <= s[index] && s[index] <= '9' {
			left := index
			for '0' <= s[index] && s[index] <= '9' {
				index++
			}
			stack = append(stack, s[left:index])
		} else if s[index] == ']' {
			var tmp []string
			for stack[len(stack)-1] != "[" {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
			}
			stack = stack[:len(stack)-1]
			times, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(strings.Join(tmp, ""), times))
			index++
		} else {
			stack = append(stack, s[index:index+1])
			index++
		}
	}
	return strings.Join(stack, "")
}

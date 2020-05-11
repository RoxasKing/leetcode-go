package codinginterviews

/*
  请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

  0 <= s 的长度 <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func replaceSpace(s string) string {
	var str string
	for i := range s {
		if s[i] == ' ' {
			str += "%20"
		} else {
			str += s[i : i+1]
		}
	}
	return str
}

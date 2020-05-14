package codinginterviews

/*
  请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"、"-1E-16"及"12e+5.4"都不是。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isNumber(s string) bool {
	s = trimSpace(s)
	if s == "" {
		return false
	}
	if !scanInteger(&s) {
		if s[0] != '.' {
			return false
		}
		s = s[1:]
		if !scanPureNumber(&s) {
			return false
		}
	} else {
		if s != "" && s[0] == '.' {
			s = s[1:]
			scanPureNumber(&s)
		}
	}
	if s != "" && s[0] == 'e' {
		s = s[1:]
		if !scanInteger(&s) {
			return false
		}
	}
	return s == ""
}

func trimSpace(s string) string {
	for s != "" && s[0] == ' ' {
		s = s[1:]
	}
	for s != "" && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	return s
}

func scanInteger(s *string) bool {
	if *s != "" && ((*s)[0] == '+' || (*s)[0] == '-') {
		*s = (*s)[1:]
	}
	return scanPureNumber(s)
}

func scanPureNumber(s *string) bool {
	old := len(*s)
	for *s != "" && '0' <= (*s)[0] && (*s)[0] <= '9' {
		*s = (*s)[1:]
	}
	return len(*s) < old
}

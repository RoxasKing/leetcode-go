package main

/*
  验证给定的字符串是否可以解释为十进制数字。
  例如:
    "0" => true
    " 0.1 " => true
    "abc" => false
    "1 a" => false
    "2e10" => true
    " -90e3   " => true
    " 1e" => false
    "e3" => false
    " 6e-1" => true
    " 99e2.5 " => false
    "53.5e93" => true
    " --6 " => false
    "-+3" => false
    "95a54e53" => false
  说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
    数字 0-9
    指数 - "e"
    正/负号 - "+"/"-"
    小数点 - "."
    当然，在输入中，这些字符的上下文也很重要。
*/

func isNumber(s string) bool {
	return isReal(trim(s))
}

func trim(s string) string {
	var i int
	for i < len(s) && s[i] == ' ' {
		i++
	}
	j := len(s) - 1
	for j >= i && s[j] == ' ' {
		j--
	}
	return s[i : j+1]
}

func isReal(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		return isNonNegReal(s[1:], false)
	}
	return isNonNegReal(s, false)
}

func isNonNegReal(s string, hasDot bool) bool {
	if len(s) == 0 {
		return false
	}
	for i, c := range s {
		switch {
		case c == '.':
			if hasDot {
				return false
			}
			if i == len(s)-1 && i != 0 {
				return true
			}
			if i+1 < len(s) && s[i+1] == 'e' {
				return i != 0 && isInteger(s[i+2:])
			}
			return isNonNegReal(s[i+1:], true)
		case c == 'e':
			if i == 0 {
				return false
			}
			return isInteger(s[i+1:])
		case c < '0' || '9' < c:
			return false
		}
	}
	return true
}

func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '-' || s[0] == '+' {
		return isNonNegativeInteger(s[1:])
	}
	return isNonNegativeInteger(s)
}

func isNonNegativeInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < '0' || '9' < c {
			return false
		}
	}
	return true
}

package main

// Difficulty:
// Hard

// Tags:
// Stack

func isValid(code string) bool {
	stk := []string{}
	top := func() int { return len(stk) - 1 }
	n := len(code)
	for i := 0; i < n; {
		if code[i] != '<' {
			if len(stk) == 0 {
				return false
			}
			i++
			continue
		}
		if i++; i == n {
			return false
		}
		if isUpperCaseLetter(code[i]) {
			j := i
			for ; j < n && isUpperCaseLetter(code[j]); j++ {
			}
			if i == j || j-i > 9 || j == n || code[j] != '>' {
				return false
			}
			stk = append(stk, code[i:j])
			i = j + 1
		} else if code[i] == '/' {
			if len(stk) == 0 {
				return false
			}
			i++
			j := i
			for ; j < n && isUpperCaseLetter(code[j]); j++ {
			}
			if i == j || j-i > 9 || j == n || code[j] != '>' || stk[top()] != code[i:j] {
				return false
			}
			if i, stk = j+1, stk[:top()]; i < n && len(stk) == 0 {
				return false
			}
		} else if code[i] == '!' {
			if len(stk) == 0 {
				return false
			}
			if i++; i+7 > n || code[i:i+7] != "[CDATA[" {
				return false
			}
			for i += 7; i+2 < n && code[i:i+3] != "]]>"; i++ {
			}
			if i += 3; i > n {
				return false
			}
		} else {
			return false
		}
	}
	return len(stk) == 0
}

func isUpperCaseLetter(c byte) bool { return 'A' <= c && c <= 'Z' }

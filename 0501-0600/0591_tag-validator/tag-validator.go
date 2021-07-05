package main

// Tags:
// Stack

func isValid(code string) bool {
	n := len(code)
	ss := StringStack{}

	for i := 0; i < n; {
		if code[i] != '<' {
			if ss.Len() == 0 {
				return false
			}
			i++
			continue
		}
		i++
		switch {
		case isUpperCaseLetter(code[i]):
			j := i
			for j < n && isUpperCaseLetter(code[j]) {
				j++
			}
			if j == n || j-i < 1 || j-i > 9 || code[j] != '>' {
				return false
			}
			ss.Push(code[i:j])
			i = j + 1
		case code[i] == '/':
			i++
			j := i
			for j < n && isUpperCaseLetter(code[j]) {
				j++
			}
			if j == n || j-i < 1 || j-i > 9 || code[j] != '>' {
				return false
			}
			if ss.Len() == 0 || ss.Pop() != code[i:j] {
				return false
			}
			i = j + 1
			if ss.Len() == 0 {
				return i == n
			}
		case code[i] == '!':
			if ss.Len() == 0 {
				return false
			}
			i++
			if i+7 > n || code[i:i+7] != "[CDATA[" {
				return false
			}
			i += 7
			for i+2 < n && code[i:i+3] != "]]>" {
				i++
			}
			if i == n {
				return false
			}
			i += 3
		default:
			return false
		}
	}

	return ss.Len() == 0
}

type StringStack []string

func (s StringStack) Len() int       { return len(s) }
func (s StringStack) Top() string    { return s[s.Len()-1] }
func (s *StringStack) Push(x string) { *s = append(*s, x) }
func (s *StringStack) Pop() string {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

func isUpperCaseLetter(letter byte) bool {
	return 'A' <= letter && letter <= 'Z'
}

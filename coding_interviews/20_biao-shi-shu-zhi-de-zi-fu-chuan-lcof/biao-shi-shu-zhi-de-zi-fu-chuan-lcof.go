package main

func isNumber(s string) bool {
	trimSpace(&s)
	if s == "" {
		return false
	}
	if !walkInteger(&s) {
		if s[0] != '.' {
			return false
		}
		s = s[1:]
		if !walkPureInteger(&s) {
			return false
		}
	} else if s != "" && s[0] == '.' {
		s = s[1:]
		walkPureInteger(&s)
	}
	if s != "" && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		if !walkInteger(&s) {
			return false
		}
	}
	return s == ""
}

func trimSpace(s *string) {
	for *s != "" && (*s)[0] == ' ' {
		*s = (*s)[1:]
	}
	for *s != "" && (*s)[len(*s)-1] == ' ' {
		*s = (*s)[:len(*s)-1]
	}
}

func walkInteger(s *string) bool {
	if *s != "" && ((*s)[0] == '-' || (*s)[0] == '+') {
		*s = (*s)[1:]
	}
	return walkPureInteger(s)
}

func walkPureInteger(s *string) bool {
	n := len(*s)
	for *s != "" && '0' <= (*s)[0] && (*s)[0] <= '9' {
		*s = (*s)[1:]
	}
	return len(*s) < n
}

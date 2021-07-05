package main

func isNumber(s string) bool {
	trimSpace(&s)
	havePrefixNum := walkInteger(&s)
	if len(s) > 0 && s[0] == '.' {
		s = s[1:]
		if ok := walkPureInteger(&s); !ok && !havePrefixNum {
			return false
		}
		havePrefixNum = true
	}
	if len(s) > 0 && (s[0] == 'E' || s[0] == 'e') {
		if !havePrefixNum {
			return false
		}
		if s = s[1:]; !walkInteger(&s) {
			return false
		}
	}
	return len(s) == 0
}

func trimSpace(s *string) {
	for len(*s) > 0 && (*s)[0] == ' ' {
		*s = (*s)[1:]
	}
	for len(*s) > 0 && (*s)[len(*s)-1] == ' ' {
		*s = (*s)[:len(*s)-1]
	}
}

func walkInteger(s *string) bool {
	if len(*s) > 0 && ((*s)[0] == '+' || (*s)[0] == '-') {
		*s = (*s)[1:]
	}
	return walkPureInteger(s)
}

func walkPureInteger(s *string) bool {
	cnt := 0
	for len(*s) > 0 && '0' <= (*s)[0] && (*s)[0] <= '9' {
		*s = (*s)[1:]
		cnt++
	}
	return cnt > 0
}

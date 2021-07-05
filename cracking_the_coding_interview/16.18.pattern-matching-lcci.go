package main

func patternMatching(pattern string, value string) bool {
	if len(pattern) == 0 {
		return value == ""
	} else if len(pattern) == 1 {
		return true
	}
	if pattern[0] == 'b' { // 重写匹配字符串，确保首个字符是 a 开头，如 bbaaa 重写为 aabbb
		bytes := []byte(pattern)
		for i := range bytes {
			if bytes[i] == 'b' {
				bytes[i] = 'a'
			} else {
				bytes[i] = 'b'
			}
		}
		pattern = string(bytes)
	}
	var countA, countB int
	for i := range pattern { // 统计匹配字符串中 a 和 b 的个数
		if pattern[i] == 'a' {
			countA++
		} else {
			countB++
		}
	}
	for i := 0; i <= len(value); i++ {
		a := value[:i]
		var offset int // 截取字符串 b 的起始偏移量
		for p := range pattern {
			if (p+1)*len(a) <= len(value) &&
				pattern[p] == 'a' &&
				value[p*len(a):(p+1)*len(a)] == a {
				offset += len(a)
			} else {
				break
			}
		}
		for j := offset; j <= len(value); j++ {
			b := value[offset:j]
			if b == a {
				continue
			}
			countLen := countA*len(a) + countB*len(b)
			if countLen == len(value) {
				var index int
				for k := range pattern {
					if pattern[k] == 'a' && value[index:index+len(a)] == a {
						index += len(a)
					} else if pattern[k] == 'b' && value[index:index+len(b)] == b {
						index += len(b)
					}
				}
				if index == len(value) {
					return true
				}
				break
			} else if countLen > len(value) {
				break
			}
		}
	}
	return false
}

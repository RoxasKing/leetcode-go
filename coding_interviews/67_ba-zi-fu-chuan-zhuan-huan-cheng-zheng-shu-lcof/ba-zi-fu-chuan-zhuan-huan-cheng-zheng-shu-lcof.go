package main

func strToInt(str string) int {
	for len(str) > 0 && str[0] == ' ' {
		str = str[1:]
	}
	if len(str) == 0 {
		return 0
	}

	sig := 1
	if str[0] == '-' {
		sig = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	out := 0
	for ; str != "" && '0' <= str[0] && str[0] <= '9'; str = str[1:] {
		if sig == 1 && (1<<31-1-int(str[0]-'0'))/10 < out {
			return 1<<31 - 1
		} else if sig == -1 && (1<<31-int(str[0]-'0'))/10 < out {
			return -1 << 31
		}
		out = out*10 + int(str[0]-'0')
	}
	out *= sig
	return out
}

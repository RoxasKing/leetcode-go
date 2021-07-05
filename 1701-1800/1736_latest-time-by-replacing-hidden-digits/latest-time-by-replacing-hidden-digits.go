package main

func maximumTime(time string) string {
	cs := []byte(time)
	if cs[0] == '?' {
		if '4' <= cs[1] && cs[1] <= '9' {
			cs[0] = '1'
		} else {
			cs[0] = '2'
		}
	}
	if cs[1] == '?' {
		if cs[0] == '2' {
			cs[1] = '3'
		} else {
			cs[1] = '9'
		}
	}
	if cs[3] == '?' {
		cs[3] = '5'
	}
	if cs[4] == '?' {
		cs[4] = '9'
	}
	return string(cs)
}

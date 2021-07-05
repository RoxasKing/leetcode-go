package main

func backspaceCompare(S string, T string) bool {
	return reBuild(S) == reBuild(T)
}

func reBuild(s string) string {
	var out string
	for i := range s {
		if s[i] != '#' {
			out += s[i : i+1]
		} else if len(out) > 0 {
			out = out[:len(out)-1]
		}
	}
	return out
}

// Two Pointers
func backspaceCompare2(S string, T string) bool {
	sPtr, tPtr := len(S)-1, len(T)-1
	sCnt, tCnt := 0, 0
	for sPtr >= 0 || tPtr >= 0 {
		for sPtr >= 0 {
			if S[sPtr] == '#' {
				sCnt++
				sPtr--
			} else if sCnt > 0 {
				sCnt--
				sPtr--
			} else {
				break
			}
		}
		for tPtr >= 0 {
			if T[tPtr] == '#' {
				tCnt++
				tPtr--
			} else if tCnt > 0 {
				tCnt--
				tPtr--
			} else {
				break
			}
		}
		if sPtr >= 0 && tPtr >= 0 {
			if S[sPtr] != T[tPtr] {
				return false
			}
		} else if sPtr >= 0 || tPtr >= 0 {
			return false
		}
		sPtr--
		tPtr--
	}
	return true
}

package main

func compress(chars []byte) int {
	i, cur, cnt := 0, chars[0], 0
	for _, ch := range chars {
		if ch == cur {
			cnt++
			continue
		}
		chars[i] = cur
		i++
		if cnt > 1 {
			tmp := []byte{}
			for ; cnt > 0; cnt /= 10 {
				tmp = append(tmp, byte(cnt%10)+'0')
			}
			for ; len(tmp) > 0; i, tmp = i+1, tmp[:len(tmp)-1] {
				chars[i] = tmp[len(tmp)-1]
			}
		}
		cur = ch
		cnt = 1
	}
	chars[i] = cur
	i++
	if cnt > 1 {
		tmp := []byte{}
		for ; cnt > 0; cnt /= 10 {
			tmp = append(tmp, byte(cnt%10)+'0')
		}
		for ; len(tmp) > 0; i, tmp = i+1, tmp[:len(tmp)-1] {
			chars[i] = tmp[len(tmp)-1]
		}
	}
	// chars = chars[:i]
	return i
}

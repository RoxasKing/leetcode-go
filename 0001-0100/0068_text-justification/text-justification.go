package main

func fullJustify(words []string, maxWidth int) []string {
	n := len(words)
	lens := make([]int, n)
	for i := range lens {
		lens[i] = countChar(words[i])
	}
	var out []string
	for l, r := 0, 0; l < n; l = r {
		count := lens[l]
		for r = l + 1; r < n && count+(r-l)+lens[r] <= maxWidth; r++ {
			count += lens[r]
		}
		if r-1-l == 0 {
			out = append(out, words[l]+makeSpace(maxWidth-count))
			continue
		}
		remain := (maxWidth - count) % (r - 1 - l)
		width := (maxWidth - count) / (r - 1 - l)
		space := makeSpace(width)
		tmp := words[l]
		for i := l + 1; i < r; i++ {
			if r == n {
				tmp += " "
			} else {
				tmp += space
				if remain > 0 {
					tmp += " "
					remain--
				}
			}
			tmp += words[i]
		}
		if r == n {
			tmp += makeSpace(maxWidth - len(tmp))
		}
		out = append(out, tmp)
	}
	return out
}

func countChar(s string) int {
	count := 0
	for s != "" {
		s = s[1:]
		count++
	}
	return count
}

func makeSpace(width int) string {
	var out string
	for i := 0; i < width; i++ {
		out += " "
	}
	return out
}

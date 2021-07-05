package main

// Tags:
// DFS + Dynamic Programming
func longestDecomposition(text string) int {
	if text == "" {
		return 0
	}
	n := len(text)
	out := 1
	for size := 1; size*2 <= n; size++ {
		if text[:size] != text[n-size:] {
			continue
		}
		i := 0
		for (i+1)*size <= n-(i+1)*size && text[i*size:(i+1)*size] == text[:size] &&
			text[i*size:(i+1)*size] == text[n-(i+1)*size:n-i*size] {
			i++
		}
		i--
		out = Max(out, 2*(i+1)+longestDecomposition(text[(i+1)*size:n-(i+1)*size]))
		break
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS + Dynamic Programming + Hash
func longestDecomposition2(text string) int {
	if text == "" {
		return 0
	}
	out := 1
	lHash, rHash := 0, 0
	n := len(text)
	l, r := 0, n-1
	base := 1
	for l < r {
		lHash = lHash*26 + int(text[l]-'a')
		rHash += base * int(text[r]-'a')
		base *= 26
		if lHash != rHash {
			l++
			r--
			continue
		}
		size := l + 1
		i := 0
		for (i+1)*size <= n-(i+1)*size && text[i*size:(i+1)*size] == text[:size] &&
			text[i*size:(i+1)*size] == text[n-(i+1)*size:n-i*size] {
			i++
		}
		i--
		out = Max(out, 2*(i+1)+longestDecomposition2(text[(i+1)*size:n-(i+1)*size]))
		break
	}
	return out
}

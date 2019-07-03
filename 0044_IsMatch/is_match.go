package main

func isMatch(s string, p string) bool {
	lens := len(s)
	lenp := len(p)

	i := -1
	for j := 0; j < lenp; j++ {
		if p[j] == '?' {
			i++
		} else if p[j] == '*' {
			if j+1 < lenp && p[j+1] == s[i] {
				i++
				j++
			} else {

			}
		} else {
			if p[j] == s[i] {
				i++
			} else {
				return false
			}
		}
	}

}

func main() {

}

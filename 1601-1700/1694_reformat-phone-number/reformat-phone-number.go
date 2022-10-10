package main

// Difficulty:
// Easy

func reformatNumber(number string) string {
	a := []rune{}
	for _, c := range number {
		if c != '-' && c != ' ' {
			a = append(a, c)
		}
	}
	o := make([]rune, 0, len(a)+len(a)/3+1)
	for i, c := range a {
		if i > 0 && i%3 == 0 && (len(a)%3 != 1 || len(a)-i > 3) || len(a)%3 == 1 && len(a)-i < 3 && i%3 == 2 {
			o = append(o, '-')
		}
		o = append(o, c)
	}
	return string(o)
}

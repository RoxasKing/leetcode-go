package main

// Tags:
// Dynamic Programming

func checkValidString(s string) bool {
	f0 := [101]bool{true}
	for i := range s {
		f1 := [101]bool{}
		for j := 0; j <= 100; j++ {
			if !f0[j] {
				continue
			}
			if s[i] == '(' && j < 100 {
				f1[j+1] = true
			}
			if s[i] == ')' && j > 0 {
				f1[j-1] = true
			}
			if s[i] == '*' {
				f1[j] = true
				if j > 0 {
					f1[j-1] = true
				}
				if j < 100 {
					f1[j+1] = true
				}
			}
		}
		f0 = f1
	}
	return f0[0]
}

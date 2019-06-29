package main

import "fmt"

func generateParenthesis(n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		var tmp []string
		if len(out) == 0 {
			tmp = append(tmp, "()")
		} else {
			for j := 0; j < len(out); j++ {
				for k := 0; k < len(out[j]); k++ {
					if k == 0 {
						tmp = append(tmp, "()"+out[j])
					} else if out[j][k-1] == '(' && out[j][k] == ')' || out[j][k-1] == '(' && out[j][k] == '(' {
						tmp = append(tmp, out[j][:k]+"()"+out[j][k:])
					} else {
						break
					}
				}
			}
		}
		out = tmp
	}
	return out
}

func main() {
	i := 3
	fmt.Println(generateParenthesis(i))
}

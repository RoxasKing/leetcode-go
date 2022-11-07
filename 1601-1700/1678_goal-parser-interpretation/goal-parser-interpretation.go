package main

import "strings"

// Difficulty:
// Easy

func interpret(command string) string {
	a := []string{}
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			a = append(a, "G")
		} else if command[i:i+2] == "()" {
			a = append(a, "o")
			i++
		} else {
			a = append(a, "al")
			i += 3
		}
	}
	return strings.Join(a, "")
}

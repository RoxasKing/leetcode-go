package main

import "strconv"

func fizzBuzz(n int) []string {
	var out []string
	for i := 1; i <= n; i++ {
		var tmp string
		if i%3 == 0 {
			tmp += "Fizz"
		}
		if i%5 == 0 {
			tmp += "Buzz"
		}
		if tmp == "" {
			tmp = strconv.Itoa(i)
		}
		out = append(out, tmp)
	}
	return out
}

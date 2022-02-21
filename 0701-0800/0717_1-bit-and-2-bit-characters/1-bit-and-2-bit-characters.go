package main

// Difficulty:
// Easy

func isOneBitCharacter(bits []int) bool {
	check := func(n int) bool {
		i := 0
		for i < n {
			if bits[i] == 1 && i+1 < n {
				i += 2
			} else if bits[i] == 0 {
				i++
			} else {
				break
			}
		}
		return i == n
	}
	return check(len(bits)) && check(len(bits)-1)
}

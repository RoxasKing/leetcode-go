package main

// Difficulty:
// Easy

// Tags:
// Simulation

func countStudents(students []int, sandwiches []int) int {
	a := 0
	for _, x := range students {
		if x == 0 {
			a++
		}
	}
	b := len(students) - a
	for _, x := range sandwiches {
		if x == 0 && a == 0 || x == 1 && b == 0 {
			break
		}
		if x == 0 {
			a--
		} else {
			b--
		}
	}
	return a + b
}

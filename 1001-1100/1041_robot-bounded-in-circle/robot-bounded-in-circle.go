package main

// Difficulty:
// Medium

// Tags:
// Simulation

func isRobotBounded(instructions string) bool {
	dir := 0 // [0:N 1:E 2:S 3:W]
	x, y := 0, 0
	for t := 4; t > 0; t-- {
		for i := range instructions {
			if instructions[i] == 'L' {
				dir = (dir + 3) % 4
			} else if instructions[i] == 'R' {
				dir = (dir + 1) % 4
			} else {
				if dir == 1 || dir == 3 {
					x += 2 - dir
				}
				if dir == 0 || dir == 2 {
					y += 1 - dir
				}
			}
		}
		if x == 0 && y == 0 {
			return true
		}
	}
	return false
}

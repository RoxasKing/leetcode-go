package main

func robot(command string, obstacles [][]int, x int, y int) bool {
	us, rs := 0, 0
	for i := range command {
		if command[i] == 'U' {
			us++
		} else {
			rs++
		}
	}

	if !reachable(command, x, y, us, rs) {
		return false
	}

	for _, ob := range obstacles {
		if ob[0] <= x && ob[1] <= y && reachable(command, ob[0], ob[1], us, rs) {
			return false
		}
	}

	return true
}

func reachable(command string, x, y int, us, rs int) bool {
	t1 := x / rs
	t2 := y / us
	base := Min(t1, t2)
	x -= base * rs
	y -= base * us

	xPos, yPos := 0, 0
	for i := range command {
		if xPos == x && yPos == y {
			return true
		}
		if command[i] == 'U' {
			yPos++
		} else {
			xPos++
		}
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

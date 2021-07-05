package main

func judgeCircle(moves string) bool {
	var row, col int
	for i := range moves {
		switch moves[i] {
		case 'U':
			row++
		case 'D':
			row--
		case 'L':
			col--
		case 'R':
			col++
		}
	}
	return row == 0 && col == 0
}

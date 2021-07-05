package main

func squareIsWhite(coordinates string) bool {
	x := int(coordinates[0] - 'a' + 1)
	y := int(coordinates[1] - '0')
	return x&1 == 1 && y&1 == 0 || x&1 == 0 && y&1 == 1
}

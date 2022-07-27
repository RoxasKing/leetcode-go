package main

// Difficulty:
// Easy

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	a, b, n := 0, 0, len(distance)
	for i := start; i != destination; i = (i + 1) % n {
		a += distance[i]
	}
	for i := destination; i != start; i = (i + 1) % n {
		b += distance[i]
	}
	if a < b {
		return a
	}
	return b
}

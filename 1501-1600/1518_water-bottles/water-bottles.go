package main

// Difficulty:
// Easy

// Tags:
// Simulation

func numWaterBottles(numBottles int, numExchange int) int {
	drinks := 0
	for empty := 0; numBottles+empty >= numExchange; empty %= numExchange {
		drinks += numBottles
		empty += numBottles
		numBottles = empty / numExchange
	}
	return drinks + numBottles
}

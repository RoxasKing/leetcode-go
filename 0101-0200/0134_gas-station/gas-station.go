package main

// Tags:
// Greedy
func canCompleteCircuit(gas []int, cost []int) int {
	var total, cur, start int
	for i := range gas {
		total += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}

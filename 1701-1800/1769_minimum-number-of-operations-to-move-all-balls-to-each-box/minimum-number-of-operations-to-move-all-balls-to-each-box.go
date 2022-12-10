package main

// Difficulty:
// Medium

// Tags:
// Prefix Sum
// Dynamic Programming

func minOperations(boxes string) []int {
	n := len(boxes)
	balls := int(boxes[n-1] - '0')
	moves := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		moves[i] = moves[i+1] + balls
		balls += int(boxes[i] - '0')
	}
	balls = 0
	move := 0
	out := make([]int, n)
	for i := range boxes {
		move += balls
		balls += int(boxes[i] - '0')
		out[i] = move + moves[i]
	}
	return out
}

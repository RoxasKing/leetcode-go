package main

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	} else if shorter == longer {
		return []int{shorter * k}
	}
	out := make([]int, k+1)
	for i := range out {
		out[i] = shorter*(k-i) + longer*i
	}
	return out
}

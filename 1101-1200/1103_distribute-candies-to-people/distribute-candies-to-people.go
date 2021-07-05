package main

func distributeCandies(candies int, numPeople int) []int {
	out := make([]int, numPeople)
	cur, idx := 1, 0
	for candies > 0 {
		if candies > cur {
			out[idx] += cur
			candies -= cur
		} else {
			out[idx] += candies
			candies = 0
		}
		cur++
		idx = (idx + 1) % numPeople
	}
	return out
}

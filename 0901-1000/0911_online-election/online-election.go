package main

// Difficulty:
// Medium

// Tags:
// Binary Search
// Hash

type TopVotedCandidate struct {
	candidates, times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(times)
	candidates := make([]int, n)
	candidates[0] = persons[0]
	votes := map[int]int{persons[0]: 1}
	for i := 1; i < n; i++ {
		votes[persons[i]]++
		if votes[persons[i]] < votes[candidates[i-1]] {
			candidates[i] = candidates[i-1]
		} else {
			candidates[i] = persons[i]
		}
	}
	return TopVotedCandidate{candidates: candidates, times: times}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.times)-1
	for l < r {
		m := (l + r) >> 1
		if this.times[m] < t {
			l = m + 1
		} else {
			r = m
		}
	}
	if this.times[l] <= t {
		return this.candidates[l]
	}
	return this.candidates[l-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

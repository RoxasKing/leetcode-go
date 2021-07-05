package main

// Tags:
// Binary Search
type TopVotedCandidate struct {
	candicate []int
	times     []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	candicate := make([]int, len(persons))
	countVote := make(map[int]int)
	candicate[0] = persons[0]
	countVote[persons[0]]++
	for i := 1; i < len(persons); i++ {
		countVote[persons[i]]++
		candicate[i] = candicate[i-1]
		if countVote[persons[i]] >= countVote[candicate[i-1]] {
			candicate[i] = persons[i]
		}
	}
	return TopVotedCandidate{
		candicate: candicate,
		times:     times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	l, r := 0, len(this.times)-1
	for l <= r {
		m := l + (r-l)>>1
		if this.times[m] < t {
			l = m + 1
		} else if this.times[m] > t {
			r = m - 1
		} else {
			return this.candicate[m]
		}
	}
	return this.candicate[l-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

package main

/*
  In an election, the i-th vote was cast for persons[i] at time times[i].

  Now, we would like to implement the following query function: TopVotedCandidate.q(int t) will return the number of the person that was leading the election at time t.

  Votes cast at time t will count towards our query.  In the case of a tie, the most recent vote (among tied candidates) wins.

  Example 1:
    Input: ["TopVotedCandidate","q","q","q","q","q","q"], [[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]
    Output: [null,0,1,1,0,0,1]
    Explanation:
      At time 3, the votes are [0], and 0 is leading.
      At time 12, the votes are [0,1,1], and 1 is leading.
      At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
      This continues for 3 more queries at time 15, 24, and 8.

  Note:
    1. 1 <= persons.length = times.length <= 5000
    2. 0 <= persons[i] <= persons.length
    3. times is a strictly increasing array with all elements in [0, 10^9].
    4. TopVotedCandidate.q is called at most 10000 times per test case.
    5. TopVotedCandidate.q(int t) is always called with t >= times[0].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/online-election
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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

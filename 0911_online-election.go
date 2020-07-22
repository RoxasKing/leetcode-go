package leetcode

/*
  在选举中，第 i 张票是在时间为 times[i] 时投给 persons[i] 的。
  现在，我们想要实现下面的查询函数： TopVotedCandidate.q(int t) 将返回在 t 时刻主导选举的候选人的编号。
  在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

  提示：
    1 <= persons.length = times.length <= 5000
    0 <= persons[i] <= persons.length
    times 是严格递增的数组，所有元素都在 [0, 10^9] 范围中。
    每个测试用例最多调用 10000 次 TopVotedCandidate.q。
    TopVotedCandidate.q(int t) 被调用时总是满足 t >= times[0]。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/online-election
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TopVotedCandidate struct {
	candicate []int
	times     []int
}

func NewTopVotedCandidate(persons []int, times []int) TopVotedCandidate {
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
		if this.times[m] <= t {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return this.candicate[l-1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

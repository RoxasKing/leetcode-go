package main

import "sort"

/*
  You are given an array of events where events[i] = [startDayi, endDayi, valuei]. The ith event starts at startDayi and ends at endDayi, and if you attend this event, you will receive a value of valuei. You are also given an integer k which represents the maximum number of events you can attend.

  You can only attend one event at a time. If you choose to attend an event, you must attend the entire event. Note that the end day is inclusive: that is, you cannot attend two events where one of them starts and the other ends on the same day.

  Return the maximum sum of values that you can receive by attending events.

  Example 1:
    Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
    Output: 7
    Explanation: Choose the green events, 0 and 1 (0-indexed) for a total value of 4 + 3 = 7.

  Example 2:
    Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
    Output: 10
    Explanation: Choose event 2 for a total value of 10.
      Notice that you cannot attend any other event as they overlap, and that you do not have to attend k events.

  Example 3:
    Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
    Output: 9
    Explanation: Although the events do not overlap, you can only attend 3 events. Pick the highest valued three.

  Constraints:
    1. 1 <= k <= events.length
    2. 1 <= k * events.length <= 10^6
    3. 1 <= startDayi <= endDayi <= 10^9
    4. 1 <= valuei <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-number-of-events-that-can-be-attended-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Discretization
// Dynamic Programming

func maxValue(events [][]int, k int) int {
	if k == 1 {
		out := 0
		for _, e := range events {
			out = Max(out, e[2])
		}
		return out
	}

	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })

	set := map[int]struct{}{}
	for _, event := range events {
		set[event[0]] = struct{}{}
		set[event[1]] = struct{}{}
	}

	pts := make([]int, 0, len(set))
	for pt := range set {
		pts = append(pts, pt)
	}
	sort.Ints(pts)

	dict := make(map[int]int)
	for i, pt := range pts {
		dict[pt] = i + 1
	}

	n := len(pts)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		copy(f[i], f[i-1])
		for ; len(events) > 0 && dict[events[0][1]] <= i; events = events[1:] {
			st, val := events[0][0], events[0][2]
			for j := 0; j < k; j++ {
				f[i][j+1] = Max(f[i][j+1], f[dict[st]-1][j]+val)
			}
		}
	}

	return f[n][k]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"sort"
)

/*
  给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。

  说明:
    1. 如果存在多种有效的行程，你可以按字符自然排序返回最小的行程组合。例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前
    2. 所有的机场都用三个大写字母表示（机场代码）。
    3. 假定所有机票至少存在一种合理的行程。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reconstruct-itinerary
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Backtracking + Hash + Graph
func findItinerary(tickets [][]string) []string {
	dict := make(map[string][]string)
	for _, ticket := range tickets {
		dict[ticket[0]] = append(dict[ticket[0]], ticket[1])
	}
	for k := range dict {
		sort.Strings(dict[k])
	}
	out := make([]string, 0, len(tickets)+1)
	out = append(out, "JFK")
	var backtrack func(string) bool
	backtrack = func(src string) bool {
		if _, ok := dict[src]; !ok || len(out) == len(tickets)+1 {
			return len(out) == len(tickets)+1
		}
		for _, dst := range dict[src] {
			out = append(out, dst)
			dict[src] = dict[src][1:]
			if backtrack(dst) {
				return true
			}
			out = out[:len(out)-1]
			dict[src] = append(dict[src], dst)
		}
		return false
	}
	_ = backtrack("JFK")
	return out
}

// DFS + Eulerian path + Hash + Graph
func findItinerary2(tickets [][]string) []string {
	dict := make(map[string][]string)
	for _, ticket := range tickets {
		dict[ticket[0]] = append(dict[ticket[0]], ticket[1])
	}
	for k := range dict {
		sort.Strings(dict[k])
	}
	out := make([]string, 0, len(tickets)+1)
	var dfs func(string)
	dfs = func(src string) {
		for len(dict[src]) != 0 {
			dst := dict[src][0]
			dict[src] = dict[src][1:]
			dfs(dst)
		}
		out = append(out, src)
	}
	dfs("JFK")
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

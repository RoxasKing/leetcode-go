package main

/*
  There are n items each belonging to zero or one of m groups where group[i] is the group that the i-th item belongs to and it's equal to -1 if the i-th item belongs to no group. The items and the groups are zero indexed. A group can have no item belonging to it.

  Return a sorted list of the items such that:
    The items that belong to the same group are next to each other in the sorted list.
	There are some relations between these items where beforeItems[i] is a list containing all the items that should come before the i-th item in the sorted array (to the left of the i-th item).

  Return any solution if there is more than one solution and return an empty list if there is no solution.

  Example 1:
    Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3,6],[],[],[]]
    Output: [6,3,4,1,5,2,0,7]

  Example 2:
    Input: n = 8, m = 2, group = [-1,-1,1,0,0,1,0,-1], beforeItems = [[],[6],[5],[6],[3],[],[4],[]]
    Output: []
    Explanation: This is the same as example 1 except that 4 needs to be before 6 in the sorted list.

  Constraints:
    1 <= m <= n <= 3 * 10^4
    group.length == beforeItems.length == n
    -1 <= group[i] <= m - 1
    0 <= beforeItems[i].length <= n - 1
    0 <= beforeItems[i][j] <= n - 1
    i != beforeItems[i][j]
    beforeItems[i] does not contain duplicates elements.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Topological Sorting + BFS
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	iEdges := make([][]int, n)
	iIndeg := make([]int, n)
	gEdges := make([][]int, m+n)
	gIndeg := make([]int, m+n)
	gItems := make([][]int, m+n)
	added := make(map[int]bool)

	gi := m
	for i := 0; i < n; i++ {
		if group[i] != -1 {
			continue
		}
		group[i] = gi
		gi++
	}

	for iv := range beforeItems {
		for _, iu := range beforeItems[iv] {
			iEdges[iu] = append(iEdges[iu], iv)
			iIndeg[iv]++

			gu, gv := group[iu], group[iv]
			idx := gu*(m+n+1) + gv
			if gu == gv || added[idx] {
				continue
			}
			added[idx] = true
			if added[gv*(m+n+1)+gu] {
				return []int{}
			}
			gEdges[gu] = append(gEdges[gu], gv)
			gIndeg[gv]++
		}
	}

	iq := []int{}
	for i := 0; i < n; i++ {
		if iIndeg[i] == 0 {
			iq = append(iq, i)
		}
	}

	cnt := 0
	for len(iq) > 0 {
		if cnt == n {
			return []int{}
		}
		cnt++
		u := iq[0]
		iq = iq[1:]
		gItems[group[u]] = append(gItems[group[u]], u)
		for _, v := range iEdges[u] {
			iIndeg[v]--
			if iIndeg[v] == 0 {
				iq = append(iq, v)
			}
		}
	}

	if cnt < n {
		return []int{}
	}

	gq := []int{}
	for i := 0; i < gi; i++ {
		if gIndeg[i] == 0 {
			gq = append(gq, i)
		}
	}

	out := make([]int, 0, n)
	for len(gq) > 0 {
		u := gq[0]
		gq = gq[1:]
		out = append(out, gItems[u]...)

		for _, v := range gEdges[u] {
			gIndeg[v]--
			if gIndeg[v] == 0 {
				gq = append(gq, v)
			}
		}
	}
	return out
}

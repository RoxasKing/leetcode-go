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
    1 <= m <= n <= 3 * 104
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

// Topological Sorting
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	itemEdges := make([][]int, n)
	itemIndeg := make([]int, n)

	groupDepends := make([][]bool, m)
	groupEdges := make([][]int, m)
	groupIndeg := make([]int, m)

	for i, items := range beforeItems {
		for _, j := range items {
			itemEdges[j] = append(itemEdges[j], i)
			itemIndeg[i]++

			groupI := group[i]
			groupJ := group[j]
			if groupI == -1 || groupJ == -1 || groupJ == groupI ||
				groupDepends[groupI] != nil && groupDepends[groupI][groupJ] {
				continue
			}

			if groupDepends[groupI] == nil {
				groupDepends[groupI] = make([]bool, m)
			}
			groupDepends[groupI][groupJ] = true

			// check groups dependeces, ifcontain circular, return empty list
			if groupDepends[groupJ] != nil && groupDepends[groupJ][groupI] {
				return []int{}
			}

			groupEdges[groupJ] = append(groupEdges[groupJ], groupI)
			groupIndeg[groupI]++
		}
	}

	itempQ := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if itemIndeg[i] == 0 {
			itempQ = append(itempQ, i)
		}
	}

	count := 0
	groupItemsAll := make([][]int, m)
	groupBeforeSelf := make([][]int, m)
	groupBeforeOther := make([][]int, m)

	for len(itempQ) > 0 {
		i := itempQ[0]
		itempQ = itempQ[1:]
		if count == n {
			return []int{}
		}
		count++
		if group[i] != -1 {
			groupItemsAll[group[i]] = append(groupItemsAll[group[i]], i)
		}
		for _, j := range itemEdges[i] {
			if group[j] != -1 {
				groupItemsAll[group[j]] = append(groupItemsAll[group[j]], j)
				if group[i] == group[j] {
					groupBeforeSelf[group[j]] = append(groupBeforeSelf[group[j]], i)
				} else {
					groupBeforeOther[group[j]] = append(groupBeforeOther[group[j]], i)
				}
			}
			itemIndeg[j]--
			if itemIndeg[j] == 0 {
				itempQ = append(itempQ, j)
			}
		}
	}

	if count != n {
		return []int{}
	}

	mark := make([]bool, n)
	out := make([]int, 0, n)

	groupQ := make([]int, 0, m)
	for g := 0; g < m; g++ {
		if groupIndeg[g] == 0 {
			groupQ = append(groupQ, g)
		}
	}

	for len(groupQ) > 0 {
		groupIdx := groupQ[0]
		groupQ = groupQ[1:]
		for _, i := range groupBeforeOther[groupIdx] {
			if !mark[i] {
				out = append(out, i)
				mark[i] = true
			}
		}
		for _, i := range groupBeforeSelf[groupIdx] {
			if !mark[i] {
				out = append(out, i)
				mark[i] = true
			}
		}
		for _, i := range groupItemsAll[groupIdx] {
			if !mark[i] {
				out = append(out, i)
				mark[i] = true
			}
		}
		for _, g := range groupEdges[groupIdx] {
			groupIndeg[g]--
			if groupIndeg[g] == 0 {
				groupQ = append(groupQ, g)
			}
		}
	}

	for i := range mark {
		if !mark[i] {
			out = append(out, i)
		}
	}

	return out
}

package main

/*
  You are given an array routes representing bus routes where routes[i] is a bus route that the ith bus repeats forever.

    For example, if routes[0] = [1, 5, 7], this means that the 0th bus travels in the sequence 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1 -> ... forever.

  You will start at the bus stop source (You are not on any bus initially), and you want to go to the bus stop target. You can travel between bus stops by buses only.

  Return the least number of buses you must take to travel from source to target. Return -1 if it is not possible.

  Example 1:
    Input: routes = [[1,2,7],[3,6,7]], source = 1, target = 6
    Output: 2
    Explanation: The best strategy is take the first bus to the bus stop 7, then take the second bus to the bus stop 6.

  Example 2:
    Input: routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target = 12
    Output: -1

  Constraints:
    1. 1 <= routes.length <= 500.
    2. 1 <= routes[i].length <= 10^5
    3. All the values of routes[i] are unique.
    4. sum(routes[i].length) <= 10^5
    5. 0 <= routes[i][j] < 10^6
    6. 0 <= source, target < 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/bus-routes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// BFS
// Hash

func numBusesToDestination(routes [][]int, source int, target int) int {
	g := map[int][]int{}
	for i, arr := range routes {
		for _, route := range arr {
			g[route] = append(g[route], i)
		}
	}

	q := [][2]int{{source, 0}}
	visited := [500]bool{}
	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		cur, times := u[0], u[1]
		for _, bus := range g[cur] {
			if visited[bus] {
				continue
			}
			visited[bus] = true
			for _, route := range routes[bus] {
				if route == cur {
					continue
				}
				if route == target {
					return times + 1
				}
				q = append(q, [2]int{route, times + 1})
			}
		}
	}
	return -1
}

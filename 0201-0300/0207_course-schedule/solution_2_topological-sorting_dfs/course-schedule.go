package main

/*
  There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

  Return true if you can finish all courses. Otherwise, return false.

  Example 1:
    Input: numCourses = 2, prerequisites = [[1,0]]
    Output: true
    Explanation: There are a total of 2 courses to take.
      To take course 1 you should have finished course 0. So it is possible.

  Example 2:
    Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
    Output: false
    Explanation: There are a total of 2 courses to take.
    To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

  Constraints:
    1. 1 <= numCourses <= 10^5
    2. 0 <= prerequisites.length <= 5000
    3. prerequisites[i].length == 2
    4. 0 <= ai, bi < numCourses
    5. All the pairs prerequisites[i] are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/course-schedule
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Topological Sorting + DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses) // 0: not visited 1: visiting 2: visited
	for _, p := range prerequisites {
		v, u := p[0], p[1]
		edges[u] = append(edges[u], v)
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && !dfs(edges, visited, i) {
			return false
		}
	}
	return true
}

func dfs(edges [][]int, visited []int, u int) bool {
	visited[u] = 1
	for _, v := range edges[u] {
		if visited[v] == 1 || visited[v] == 0 && !dfs(edges, visited, v) {
			return false
		}
	}
	visited[u] = 2
	return true
}

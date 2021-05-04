package main

/*
  There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

    For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

  Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

  Example 1:
    Input: numCourses = 2, prerequisites = [[1,0]]
    Output: [0,1]
    Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

  Example 2:
    Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
    Output: [0,2,1,3]
    Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
    So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

  Example 3:
    Input: numCourses = 1, prerequisites = []
    Output: [0]

  Constraints:
    1. 1 <= numCourses <= 2000
    2. 0 <= prerequisites.length <= numCourses * (numCourses - 1)
    3. prerequisites[i].length == 2
    4. 0 <= ai, bi < numCourses
    5. ai != bi
    6. All the pairs [ai, bi] are distinct.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/course-schedule-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Topological Sorting + BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indeg[p[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	out := make([]int, 0, numCourses)
	for len(q) > 0 {
		if len(out) == numCourses {
			return nil
		}
		c := q[0]
		q = q[1:]
		out = append(out, c)
		for _, nc := range edges[c] {
			indeg[nc]--
			if indeg[nc] == 0 {
				q = append(q, nc)
			}
		}
	}

	if len(out) < numCourses {
		return nil
	}
	return out
}

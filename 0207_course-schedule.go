package leetcode

/*
  你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
  在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
  给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

  提示：
    输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
    你可以假定输入的先决条件中没有重复的边。
    1 <= numCourses <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/course-schedule
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Topological Sorting + BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		queue  = make([]int, 0, numCourses)
		count  int
		course int
	)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indeg[p[0]]++
	}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) != 0 {
		course, queue = queue[0], queue[1:]
		if count == numCourses {
			return false
		}
		count++
		for _, c := range edges[course] {
			indeg[c]--
			if indeg[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	return count == numCourses
}

// Topological Sorting + DFS
func canFinish2(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	can := true
	var dfs func(int)
	dfs = func(course int) {
		visited[course] = 1
		for _, c := range edges[course] {
			if visited[c] == 0 {
				dfs(c)
				if !can {
					return
				}
			} else if visited[c] == 1 {
				can = false
				return
			}
		}
		visited[course] = 2
	}
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	for i := 0; i < numCourses && can; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return can
}

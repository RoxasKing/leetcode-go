package leetcode

/*
  现在你总共有 n 门课需要选，记为 0 到 n-1。
  在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]
  给定课程总量以及它们的先决条件，返回你为了学完所有课程所安排的学习顺序。
  可能会有多个正确的顺序，你只要返回一种就可以了。如果不可能完成所有课程，返回一个空数组。

  说明:
    输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
    你可以假定输入的先决条件中没有重复的边。
  提示:
    这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
    通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
    拓扑排序也可以通过 BFS 完成。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/course-schedule-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Topological Sorting + DFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	out := make([]int, 0, numCourses)
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	var invalid bool
	var dfs func(int)
	dfs = func(course int) {
		visited[course] = 1
		for _, c := range edges[course] {
			if visited[c] == 0 {
				dfs(c)
				if invalid {
					return
				}
			} else if visited[c] == 1 {
				invalid = true
				return
			}
		}
		visited[course] = 2
		out = append(out, course)
	}
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	for i := 0; i < numCourses && !invalid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if invalid {
		return nil
	}
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

// Topological Sorting + BFS
func findOrder2(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indeg[p[0]]++
	}
	var queue []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	if len(queue) == 0 {
		return nil
	}
	var index int
	out := make([]int, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		if index < numCourses {
			out[index] = course
			index++
		} else {
			return nil
		}
		for _, c := range edges[course] {
			indeg[c]--
			if indeg[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	if index < numCourses {
		return nil
	}
	return out
}

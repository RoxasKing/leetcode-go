package main

// Tags:
// DFS

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	dict := make(map[int]int)
	for i, e := range employees {
		dict[e.Id] = i
	}
	return dfs(employees, dict, id)
}

func dfs(employees []*Employee, dict map[int]int, id int) int {
	out := employees[dict[id]].Importance
	for _, id := range employees[dict[id]].Subordinates {
		out += dfs(employees, dict, id)
	}
	return out
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

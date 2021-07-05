package main

// Tags:
// DFS + Recursion
func canVisitAllRooms(rooms [][]int) bool {
	mark := make([]bool, len(rooms))
	var dfs func(int)
	dfs = func(index int) {
		if mark[index] {
			return
		}
		mark[index] = true
		for _, key := range rooms[index] {
			dfs(key)
		}
	}
	dfs(0)
	for i := range mark {
		if !mark[i] {
			return false
		}
	}
	return true
}

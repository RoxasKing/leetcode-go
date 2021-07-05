package main

// Tags:
// Hash
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	cnt := make(map[int]map[int]bool)
	for _, log := range logs {
		if cnt[log[0]] == nil {
			cnt[log[0]] = make(map[int]bool)
		}
		cnt[log[0]][log[1]] = true
	}
	out := make([]int, k)
	for _, v := range cnt {
		out[len(v)-1]++
	}
	return out
}

package main

// Difficulty:
// Hard

// Tags:
// Hash
// BFS

var rs = []int{0, 1, 0, -1}
var cs = []int{1, 0, -1, 0}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	max := len(blocked) * len(blocked) / 2
	ban := map[[2]int]struct{}{}
	for _, co := range blocked {
		x, y := co[0], co[1]
		ban[[2]int{x, y}] = struct{}{}
	}

	check := func(source, target []int) (bool, bool) {
		cnt := 1
		vis := map[[2]int]struct{}{}
		vis[[2]int{source[0], source[1]}] = struct{}{}
	LOOP:
		for q := [][]int{source}; len(q) > 0; q = q[1:] {
			co := q[0]
			xi, yi := co[0], co[1]
			for i := 0; i < 4; i++ {
				xj, yj := xi+rs[i], yi+cs[i]
				if xj < 0 || 1e6-1 < xj || yj < 0 || 1e6-1 < yj {
					continue
				}
				if xj == target[0] && yj == target[1] {
					return true, true
				}
				if _, ok := ban[[2]int{xj, yj}]; ok {
					continue
				}
				if _, ok := vis[[2]int{xj, yj}]; ok {
					continue
				}
				vis[[2]int{xj, yj}] = struct{}{}
				q = append(q, []int{xj, yj})
				if cnt++; cnt > max {
					break LOOP
				}
			}
		}
		return cnt > max, false
	}

	if ok1, ok2 := check(source, target); ok2 {
		return true
	} else if !ok1 {
		return false
	}

	ok, _ := check(target, source)
	return ok
}

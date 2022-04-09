package main

// Difficulty:
// Hard

// Tags:
// Math

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for (sx < tx || sy < ty) && tx != ty {
		if tx < ty {
			if res := sy + (ty-sy)%tx; res == ty {
				break
			} else {
				ty = res
			}
		} else {
			if res := sx + (tx-sx)%ty; res == tx {
				break
			} else {
				tx = res
			}
		}
	}
	return sx == tx && sy == ty
}

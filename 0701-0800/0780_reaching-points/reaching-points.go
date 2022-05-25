package main

// Difficulty:
// Hard

// Tags:
// Math

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for (sx < tx || sy < ty) && tx != ty {
		if tx < ty {
			t := (ty-sy)%tx + sy
			if t == ty {
				return false
			}
			ty = t
		} else {
			t := (tx-sx)%ty + sx
			if t == tx {
				return false
			}
			tx = t
		}
	}
	return sx == tx && sy == ty
}

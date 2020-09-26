package main

/*
  给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
*/

func maxPoints(points [][]int) int {
	pointCount := make(map[[2]int]int)
	for _, point := range points {
		pointCount[[2]int{point[0], point[1]}]++
	}
	n := len(pointCount)
	if n < 3 {
		return len(points)
	}
	pointSet := make([][3]int, 0, n)
	for point, count := range pointCount {
		pointSet = append(pointSet, [3]int{point[0], point[1], count})
	}
	var max int
	for i, point := range pointSet {
		x0, y0, count0 := point[0], point[1], point[2]
		slopeCount := make(map[float64]int)
		pointWithSameXCount := 0
		for j := 0; j < i; j++ {
			x, y, count := pointSet[j][0], pointSet[j][1], pointSet[j][2]
			if x0 == x {
				pointWithSameXCount += count
				max = Max(max, pointWithSameXCount+count0)
			} else {
				slope := float64(y0-y) / float64(x0-x)
				slopeCount[slope] += count
				max = Max(max, slopeCount[slope]+count0)
			}
		}
	}
	return max
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxPoints2(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}
	var max int
	for i := 0; i < n; i++ {
		var cur, countCenter int
		dict := make(map[[2]int]int)
		for j := i + 1; j < n; j++ {
			a, b := points[j][0]-points[i][0], points[j][1]-points[i][1]
			if a == 0 && b == 0 {
				countCenter++
				continue
			}
			div := Gcd(a, b)
			key := [2]int{a / div, b / div}
			dict[key]++
			cur = Max(cur, dict[key])
		}
		max = Max(max, cur+countCenter+1)
	}
	return max
}

// Gcd: the greatest common divisor of a and b
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

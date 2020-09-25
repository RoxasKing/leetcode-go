package main

/*
  给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
*/

func maxPoints(points [][]int) int {
	pointMap := map[[2]int]int{}
	for _, p := range points {
		pointMap[[2]int{p[0], p[1]}]++
	}
	n := len(pointMap)
	if n <= 2 {
		return len(points)
	}
	pointCount := make([][3]int, n)
	var index int
	for point, count := range pointMap {
		pointCount[index] = [3]int{point[0], point[1], count}
		index++
	}
	var max int
	for i, p := range pointCount {
		x, y := p[0], p[1]
		var dict = make(map[float64]int)
		var pointWithSameX int
		for j := 0; j < i; j++ {
			if x == pointCount[j][0] {
				pointWithSameX += pointCount[j][2]
				max = Max(max, pointWithSameX+p[2])
			} else {
				k := float64(y-pointCount[j][1]) / float64(x-pointCount[j][0])
				dict[k] += pointCount[j][2]
				max = Max(max, dict[k]+p[2])
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

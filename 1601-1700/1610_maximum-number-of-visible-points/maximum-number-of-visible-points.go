package main

import (
	"math"
	"sort"
)

// Tags:
// Two Pointerss
func visiblePoints(points [][]int, angle int, location []int) int {
	duplicate := 0
	angles := []float64{}
	for _, p := range points {
		x, y := p[0]-location[0], p[1]-location[1]
		if x == 0 && y == 0 {
			duplicate++
			continue
		}
		angle := math.Atan2(float64(y), float64(x)) * 180 / math.Pi
		angles = append(angles, angle)
	}
	sort.Float64s(angles)
	for _, angle := range angles {
		angles = append(angles, angle+360)
	}
	max := 0
	for l, r := 0, 0; l < len(angles); l++ {
		for r < len(angles) && (angles[r]-angles[l])-float64(angle) < 1e-8 {
			r++
		}
		if r-l > max {
			max = r - l
		}
	}
	return max + duplicate
}

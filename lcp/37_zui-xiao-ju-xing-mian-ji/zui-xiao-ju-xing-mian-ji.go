package main

import (
	"math"
	"runtime/debug"
	"sort"
)

// Tags:
// Sort
// Greedy

func minRecSize(lines [][]int) float64 {
	sort.Slice(lines, func(i, j int) bool {
		if lines[i][0] != lines[j][0] {
			return lines[i][0] > lines[j][0]
		}
		return lines[i][1] > lines[j][1]
	})

	xMax, xMin := -1e9, 1e9
	yMax, yMin := -1e9, 1e9

	pk, pbMax, pbMin := -1, 0, 0
	for i := 0; i < len(lines); i++ {
		ck, cbMax, cbMin := lines[i][0], lines[i][1], lines[i][1]
		for ; i+1 < len(lines) && lines[i][0] == lines[i+1][0]; i++ {
			cbMin = lines[i+1][1]
		}
		if pk != -1 {
			var x, y float64
			x = float64(cbMax-pbMin) / float64(pk-ck)
			y = float64(ck)*x + float64(cbMax)
			xMin = math.Min(xMin, x)
			yMin = math.Min(yMin, y)
			xMax = math.Max(xMax, x)
			yMax = math.Max(yMax, y)
			x = float64(cbMin-pbMax) / float64(pk-ck)
			y = float64(ck)*x + float64(cbMin)
			xMin = math.Min(xMin, x)
			yMin = math.Min(yMin, y)
			xMax = math.Max(xMax, x)
			yMax = math.Max(yMax, y)
		}
		pk, pbMax, pbMin = ck, cbMax, cbMin
	}

	if xMax <= xMin || yMax <= yMin {
		return 0
	}
	return (xMax - xMin) * (yMax - yMin)
}

func init() { debug.SetGCPercent(-1) }

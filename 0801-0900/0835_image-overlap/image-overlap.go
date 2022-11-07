package main

// Difficulty:
// Medium

func largestOverlap(img1 [][]int, img2 [][]int) int {
	o, n := 0, len(img1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c1, c2, c3, c4 := 0, 0, 0, 0
			for x := 0; x+i < n; x++ {
				for y := 0; y+j < n; y++ {
					if img1[x][y] == 1 && img1[x][y] == img2[x+i][y+j] {
						c1++
					}
					if img2[x][y] == 1 && img1[x+i][y+j] == img2[x][y] {
						c2++
					}
					if img1[x+i][y] == 1 && img1[x+i][y] == img2[x][y+j] {
						c3++
					}
					if img2[x+i][y] == 1 && img1[x][y+j] == img2[x+i][y] {
						c4++
					}
				}
			}
			o = max(o, max(max(max(c1, c2), c3), c4))
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

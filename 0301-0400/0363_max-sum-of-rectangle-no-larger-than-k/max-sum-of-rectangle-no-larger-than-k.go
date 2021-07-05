package main

func maxSumSubmatrix(matrix [][]int, k int) int {
	out := -1 << 31
	m, n := len(matrix), len(matrix[0])
	for u := 0; u < m; u++ { // upper boundary
		sumC := make([]int, n)
		for i := u; i < m; i++ { // lower boundary
			for j := 0; j < n; j++ {
				sumC[j] += matrix[i][j]
			}

			sum := sumC[0]
			maxS := sum // max sum of current rectangle
			for j := 1; j < n; j++ {
				if sum > 0 {
					sum += sumC[j]
				} else {
					sum = sumC[j]
				}
				maxS = Max(maxS, sum)
			}
			if maxS < k {
				out = Max(out, maxS)
				continue
			} else if maxS == k {
				return k
			}

			for l := 0; l < n; l++ { // left boundary
				sum := 0
				for j := l; j < n; j++ { // right boundary
					sum += sumC[j]
					if sum < k {
						out = Max(out, sum)
					} else if sum == k {
						return k
					}
				}
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

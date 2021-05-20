package main

import (
	"math/rand"
)

/*
  You are given a 2D matrix of size m x n, consisting of non-negative integers. You are also given an integer k.

  The value of coordinate (a, b) of the matrix is the XOR of all matrix[i][j] where 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).

  Find the kth largest value (1-indexed) of all the coordinates of matrix.

  Example 1:
    Input: matrix = [[5,2],[1,6]], k = 1
    Output: 7
    Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.

  Example 2:
    Input: matrix = [[5,2],[1,6]], k = 2
    Output: 5
    Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.

  Example 3:
    Input: matrix = [[5,2],[1,6]], k = 3
    Output: 4
    Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.

  Example 4:
    Input: matrix = [[5,2],[1,6]], k = 4
    Output: 0
    Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which is the 4th largest value.

  Constraints:
    1. m == matrix.length
    2. n == matrix[i].length
    3. 1 <= m, n <= 1000
    4. 0 <= matrix[i][j] <= 10^6
    5. 1 <= k <= m * n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-kth-largest-xor-coordinate-value
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Bit Manipulation
// Prefix Sum
// Quick Select

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	arr := make([]int, 0, m*n)
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		f[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			f[i+1][j+1] = matrix[i][j] ^ f[i][j+1] ^ f[i+1][j] ^ f[i][j]
			arr = append(arr, f[i+1][j+1])
		}
	}

	return quickSelect(arr, m*n-k)
}

func quickSelect(arr []int, k int) int {
	for l, r := 0, len(arr)-1; l < r; {
		pivotIdx := rand.Intn(r+1-l) + l
		arr[pivotIdx], arr[r] = arr[r], arr[pivotIdx]
		i, j := l-1, r
		for {
			for i++; i < r && arr[i] < arr[r]; i++ {
			}
			for j--; j > l && arr[j] > arr[r]; j-- {
			}
			if i >= j {
				break
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		arr[i], arr[r] = arr[r], arr[i]
		if i < k {
			l = i + 1
		} else if i > k {
			r = i - 1
		} else {
			break
		}
	}
	return arr[k]
}

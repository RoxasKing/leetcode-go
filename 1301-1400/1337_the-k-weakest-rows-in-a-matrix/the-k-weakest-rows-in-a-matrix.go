package main

import "sort"

/*
  You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of all the 0's in each row.

  A row i is weaker than a row j if one of the following is true:

    1. The number of soldiers in row i is less than the number of soldiers in row j.
    2. Both rows have the same number of soldiers and i < j.

  Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.

  Example 1:
    Input: mat =
           [[1,1,0,0,0],
            [1,1,1,1,0],
            [1,0,0,0,0],
            [1,1,0,0,0],
            [1,1,1,1,1]],
           k = 3
    Output: [2,0,3]
    Explanation:
      The number of soldiers in each row is:
      - Row 0: 2
      - Row 1: 4
      - Row 2: 1
      - Row 3: 2
      - Row 4: 5
      The rows ordered from weakest to strongest are [2,0,3,1,4].

  Example 2:
    Input: mat =
           [[1,0,0,0],
            [1,1,1,1],
            [1,0,0,0],
            [1,0,0,0]],
           k = 2
    Output: [0,2]
    Explanation:
      The number of soldiers in each row is:
      - Row 0: 1
      - Row 1: 4
      - Row 2: 1
      - Row 3: 1
      The rows ordered from weakest to strongest are [0,2,3,1].

  Constraints:
    1. m == mat.length
    2. n == mat[i].length
    3. 2 <= n, m <= 100
    4. 1 <= k <= m
    5. matrix[i][j] is either 0 or 1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func kWeakestRows(mat [][]int, k int) []int {
	n, m := len(mat), len(mat[0])
	arr := make([]int, n)
	for i := range arr {
		arr[i] = sort.Search(m, func(j int) bool { return mat[i][j] == 0 })
	}

	out := make([]int, n)
	for i := range out {
		out[i] = i
	}

	sort.Slice(out, func(i, j int) bool {
		if arr[out[i]] != arr[out[j]] {
			return arr[out[i]] < arr[out[j]]
		}
		return out[i] < out[j]
	})

	return out[:k]
}

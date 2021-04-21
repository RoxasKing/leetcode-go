package main

/*
  An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).

  Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.

  To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.

  At the end, return the modified image.

  Example 1:
    Input:
      image = [[1,1,1],[1,1,0],[1,0,1]]
      sr = 1, sc = 1, newColor = 2
    Output: [[2,2,2],[2,2,0],[2,0,1]]
    Explanation:
      From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
      by a path of the same color as the starting pixel are colored with the new color.
      Note the bottom corner is not colored 2, because it is not 4-directionally connected
      to the starting pixel.

  Note:
    1. The length of image and image[0] will be in the range [1, 50].
    2. The given starting pixel will satisfy 0 <= sr < image.length and 0 <= sc < image[0].length.
    3. The value of each color in image[i][j] and newColor will be an integer in [0, 65535].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/flood-fill
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(image), len(image[0])
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
		copy(out[i], image[i])
	}
	out[sr][sc] = newColor
	q := [][]int{{sr, sc}}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		r, c := e[0], e[1]
		for _, f := range forwards {
			nr, nc := r+f[0], c+f[1]
			if nr < 0 || m-1 < nr || nc < 0 || n-1 < nc || image[nr][nc] != image[r][c] || out[nr][nc] == newColor {
				continue
			}
			out[nr][nc] = newColor
			q = append(q, []int{nr, nc})
		}
	}
	return out
}

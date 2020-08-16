package leetcode

/*
  有一幅以二维整数数组表示的图画，每一个整数表示该图画的像素值大小，数值在 0 到 65535 之间。
  给你一个坐标 (sr, sc) 表示图像渲染开始的像素值（行 ，列）和一个新的颜色值 newColor，让你重新上色这幅图像。
  为了完成上色工作，从初始坐标开始，记录初始坐标的上下左右四个方向上像素值与初始坐标相同的相连像素点，接着再记录这四个方向上符合条件的像素点与他们对应四个方向上像素值与初始坐标相同的相连像素点，……，重复该过程。将所有有记录的像素点的颜色值改为新的颜色值。
  最后返回经过上色渲染后的图像。

  注意:
    image 和 image[0] 的长度在范围 [1, 50] 内。
    给出的初始点将满足 0 <= sr < image.length 和 0 <= sc < image[0].length。
    image[i][j] 和 newColor 表示的颜色值在范围 [0, 65535]内。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/flood-fill
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 || image[sr][sc] == newColor {
		return image
	}
	newImage := make([][]int, len(image))
	for i := range newImage {
		newImage[i] = make([]int, len(image[0]))
		copy(newImage[i], image[i])
	}
	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][]int{{sr, sc}}
	for len(queue) != 0 {
		row, col := queue[0][0], queue[0][1]
		newImage[row][col] = newColor
		queue = queue[1:]
		for _, move := range moves {
			newRow, newCol := row+move[0], col+move[1]
			if 0 <= newRow && newRow < len(image) && 0 <= newCol && newCol < len(image[0]) &&
				newImage[newRow][newCol] == image[sr][sc] {
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
	return newImage
}

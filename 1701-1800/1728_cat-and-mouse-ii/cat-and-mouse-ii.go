package main

/*
  A game is played by a cat and a mouse named Cat and Mouse.

  The environment is represented by a grid of size rows x cols, where each element is a wall, floor, player (Cat, Mouse), or food.

    1. Players are represented by the characters 'C'(Cat),'M'(Mouse).
    2. Floors are represented by the character '.' and can be walked on.
    3. Walls are represented by the character '#' and cannot be walked on.
    4. Food is represented by the character 'F' and can be walked on.
    5. There is only one of each character 'C', 'M', and 'F' in grid.

  Mouse and Cat play according to the following rules:

    1. Mouse moves first, then they take turns to move.
    2. During each turn, Cat and Mouse can jump in one of the four directions (left, right, up, down). They cannot jump over the wall nor outside of the grid.
    3. catJump, mouseJump are the maximum lengths Cat and Mouse can jump at a time, respectively. Cat and Mouse can jump less than the maximum length.
    4. Staying in the same position is allowed.
    5. Mouse can jump over Cat.

  The game can end in 4 ways:

    1. If Cat occupies the same position as Mouse, Cat wins.
    2. If Cat reaches the food first, Cat wins.
    3. If Mouse reaches the food first, Mouse wins.
    4. If Mouse cannot get to the food within 1000 turns, Cat wins.

  Given a rows x cols matrix grid and two integers catJump and mouseJump, return true if Mouse can win the game if both Cat and Mouse play optimally, otherwise return false.

  Example 1:
    Input: grid = ["####F","#C...","M...."], catJump = 1, mouseJump = 2
    Output: true
    Explanation: Cat cannot catch Mouse on its turn nor can it get the food before Mouse.

  Example 2:
    Input: grid = ["M.C...F"], catJump = 1, mouseJump = 4
    Output: true

  Example 3:
    Input: grid = ["M.C...F"], catJump = 1, mouseJump = 3
    Output: false

  Example 4:
    Input: grid = ["C...#","...#F","....#","M...."], catJump = 2, mouseJump = 5
    Output: false

  Example 5:
    Input: grid = [".M...","..#..","#..#.","C#.#.","...#F"], catJump = 3, mouseJump = 1
    Output: true

  Constraints:
    1. rows == grid.length
    2. cols = grid[i].length
    3. 1 <= rows, cols <= 8
    4. grid[i][j] consist only of characters 'C', 'M', 'F', '.', and '#'.
    5. There is only one of each character 'C', 'M', and 'F' in grid.
    6. 1 <= catJump, mouseJump <= 8

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/cat-and-mouse-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// BFS
// Hash

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	rows, cols := len(grid), len(grid[0])
	var x1, y1, x2, y2 int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'M' {
				x1, y1 = i, j
			} else if grid[i][j] == 'C' {
				x2, y2 = i, j
			}
		}
	}

	f := map[[5]int]int{}
	return dp(f, grid, catJump, mouseJump, rows, cols, x1, y1, x2, y2, 0) == 1
}

func dp(f map[[5]int]int, grid []string, catJump, mouseJump, m, n, x1, y1, x2, y2, k int) int {
	if k >= 70 {
		return 0
	}
	if val, ok := f[[5]int{x1, y1, x2, y2, k}]; ok {
		return val
	}
	if k%2 == 0 {
		for _, d := range direction {
			for j := 0; j <= mouseJump; j++ {
				x, y := x1+j*d[0], y1+j*d[1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
					break
				} else if x == x2 && y == y2 {
					continue
				}
				if grid[x][y] == 'F' || dp(f, grid, catJump, mouseJump, m, n, x, y, x2, y2, k+1) == 1 {
					f[[5]int{x1, y1, x2, y2, k}] = 1
					return 1
				}
			}
		}
		f[[5]int{x1, y1, x2, y2, k}] = 0
		return 0
	} else {
		for _, d := range direction {
			for j := 0; j <= catJump; j++ {
				x, y := x2+j*d[0], y2+j*d[1]
				if x < 0 || m-1 < x || y < 0 || n-1 < y || grid[x][y] == '#' {
					break
				}
				if x == x1 && y == y1 || grid[x][y] == 'F' ||
					dp(f, grid, catJump, mouseJump, m, n, x1, y1, x, y, k+1) == 0 {
					f[[5]int{x1, y1, x2, y2, k}] = 0
					return 0
				}
			}
		}
		f[[5]int{x1, y1, x2, y2, k}] = 1
		return 1
	}
}

var direction = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

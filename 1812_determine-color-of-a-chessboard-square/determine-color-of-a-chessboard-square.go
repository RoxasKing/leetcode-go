package main

/*
  You are given coordinates, a string that represents the coordinates of a square of the chessboard. Below is a chessboard for your reference.

  Return true if the square is white, and false if the square is black.

  The coordinate will always represent a valid chessboard square. The coordinate will always have the letter first, and the number second.

  Example 1:
    Input: coordinates = "a1"
    Output: false
    Explanation: From the chessboard above, the square with coordinates "a1" is black, so return false.

  Example 2:
    Input: coordinates = "h3"
    Output: true
    Explanation: From the chessboard above, the square with coordinates "h3" is white, so return true.

  Example 3:
    Input: coordinates = "c7"
    Output: false

  Constraints:
    1. coordinates.length == 2
    2. 'a' <= coordinates[0] <= 'h'
    3. '1' <= coordinates[1] <= '8'

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/determine-color-of-a-chessboard-square
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func squareIsWhite(coordinates string) bool {
	x := int(coordinates[0] - 'a' + 1)
	y := int(coordinates[1] - '0')
	return x&1 == 1 && y&1 == 0 || x&1 == 0 && y&1 == 1
}

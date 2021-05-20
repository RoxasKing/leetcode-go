package main

/*
  Three stones are on a number line at positions a, b, and c.

  Each turn, you pick up a stone at an endpoint (ie., either the lowest or highest position stone), and move it to an unoccupied position between those endpoints.  Formally, let's say the stones are currently at positions x, y, z with x < y < z.  You pick up the stone at either position x or position z, and move that stone to an integer position k, with x < k < z and k != y.

  The game ends when you cannot make any more moves, ie. the stones are in consecutive positions.

  When the game ends, what is the minimum and maximum number of moves that you could have made?  Return the answer as an length 2 array: answer = [minimum_moves, maximum_moves]

  Example 1:
    Input: a = 1, b = 2, c = 5
    Output: [1,2]
    Explanation: Move the stone from 5 to 3, or move the stone from 5 to 4 to 3.

  Example 2:
    Input: a = 4, b = 3, c = 2
    Output: [0,0]
    Explanation: We cannot make any moves.

  Example 3:
    Input: a = 3, b = 5, c = 1
    Output: [1,2]
    Explanation: Move the stone from 1 to 4; or move the stone from 1 to 2 to 4.

  Note:
    1. 1 <= a <= 100
    2. 1 <= b <= 100
    3. 1 <= c <= 100
    4. a != b, b != c, c != a

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/moving-stones-until-consecutive
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Brain Teaser

func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}

	out := []int{0, 0}

	if a+2 == c {
		return out
	}

	if a+2 >= b || b+2 >= c {
		out[0] = 1
	} else {
		out[0] = 2
	}

	out[1] += c - a - 2

	return out
}

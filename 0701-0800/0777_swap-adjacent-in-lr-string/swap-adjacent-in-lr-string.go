package main

/*
  In a string composed of 'L', 'R', and 'X' characters, like "RXXLRXRXL", a move consists of either replacing one occurrence of "XL" with "LX", or replacing one occurrence of "RX" with "XR". Given the starting string start and the ending string end, return True if and only if there exists a sequence of moves to transform one string to the other.

  Example 1:
    Input: start = "RXXLRXRXL", end = "XRLXXRRLX"
    Output: true
    Explanation: We can transform start to end following these steps:
      RXXLRXRXL ->
      XRXLRXRXL ->
      XRLXRXRXL ->
      XRLXXRRXL ->
      XRLXXRRLX

  Example 2:
    Input: start = "X", end = "L"
    Output: false

  Example 3:
    Input: start = "LLR", end = "RRL"
    Output: false

  Example 4:
    Input: start = "XL", end = "LX"
    Output: true

  Example 5:
    Input: start = "XLLR", end = "LXLX"
    Output: false

  Constraints:
    1 <= start.length <= 10^4
    start.length == end.length
    Both start and end will only consist of characters in 'L', 'R', and 'X'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/swap-adjacent-in-lr-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Brain Teaser

func canTransform(start string, end string) bool {
	for i, j, n := -1, -1, len(start); ; {
		for i++; i < n && start[i] == 'X'; i++ {
		}
		for j++; j < n && end[j] == 'X'; j++ {
		}

		if i == n && j == n {
			break
		}

		if i == n || j == n || start[i] != end[j] || start[i] == 'L' && i < j || start[i] == 'R' && i > j {
			return false
		}
	}
	return true
}

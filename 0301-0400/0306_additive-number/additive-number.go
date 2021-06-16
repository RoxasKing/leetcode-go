package main

/*
  Additive number is a string whose digits can form additive sequence.

  A valid additive sequence should contain at least three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

  Given a string containing only digits '0'-'9', write a function to determine if it's an additive number.

  Note: Numbers in the additive sequence cannot have leading zeros, so sequence 1, 2, 03 or 1, 02, 3 is invalid.

  Example 1:
    Input: "112358"
    Output: true
    Explanation: The digits can form an additive sequence: 1, 1, 2, 3, 5, 8.
                 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

  Example 2:
    Input: "199100199"
    Output: true
    Explanation: The additive sequence is: 1, 99, 100, 199.
                 1 + 99 = 100, 99 + 100 = 199

  Constraints:
    num consists only of digits '0'-'9'.
    1 <= num.length <= 35

  Follow up:
    How would you handle overflow for very large input integers?


  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/additive-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking

func isAdditiveNumber(num string) bool {
	return dfs(num, len(num), 0, -1, -1, 0)
}

func dfs(num string, n, i, p1, p2, cur int) bool {
	if i == n {
		return p1 != -1 && p2 != -1 && p1+p2 == cur
	}

	cur = cur*10 + int(num[i]-'0')

	if cur != 0 && dfs(num, n, i+1, p1, p2, cur) {
		return true
	}

	if p1 == -1 {
		return dfs(num, n, i+1, cur, p2, 0)
	}

	if p2 == -1 {
		return dfs(num, n, i+1, p1, cur, 0)
	}

	if p1+p2 == cur {
		return dfs(num, n, i+1, p2, cur, 0)
	}

	return false
}

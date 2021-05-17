package main

/*
  Given a binary string s, return the minimum number of character swaps to make it alternating, or -1 if it is impossible.

  The string is called alternating if no two adjacent characters are equal. For example, the strings "010" and "1010" are alternating, while the string "0100" is not.

  Any two characters may be swapped, even if they are not adjacent.

  Example 1:
    Input: s = "111000"
    Output: 1
    Explanation: Swap positions 1 and 4: "111000" -> "101010"
      The string is now alternating.

  Example 2:
    Input: s = "010"
    Output: 0
    Explanation: The string is already alternating, no swaps are needed.

  Example 3:
    Input: s = "1110"
    Output: -1

  Constraints:
    1. 1 <= s.length <= 1000
    2. s[i] is either '0' or '1'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math + Greedy Algorithm

func minSwaps(s string) int {
	cnt := 0
	for _, ch := range s {
		if ch == '0' {
			cnt++
		} else {
			cnt--
		}
	}

	if Abs(cnt) > 1 {
		return -1
	}

	cur := 0
	if cnt < 0 {
		cur = 1
	}
	diff := 0
	for i := range s {
		if int(s[i]-'0') != cur {
			diff++
		}
		cur ^= 1
	}
	diff >>= 1

	if cnt != 0 {
		return diff
	}

	out := diff
	cur, diff = 1, 0
	for i := range s {
		if int(s[i]-'0') != cur {
			diff++
		}
		cur ^= 1
	}
	diff >>= 1

	out = Min(out, diff)

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

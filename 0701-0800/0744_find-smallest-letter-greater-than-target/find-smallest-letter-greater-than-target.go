package main

/*
  Given a list of sorted characters letters containing only lowercase letters, and given a target letter target, find the smallest element in the list that is larger than the given target.

  Letters also wrap around. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

  Examples:
    Input:
    letters = ["c", "f", "j"]
    target = "a"
    Output: "c"

    Input:
    letters = ["c", "f", "j"]
    target = "c"
    Output: "f"

    Input:
    letters = ["c", "f", "j"]
    target = "d"
    Output: "f"

    Input:
    letters = ["c", "f", "j"]
    target = "g"
    Output: "j"

    Input:
    letters = ["c", "f", "j"]
    target = "j"
    Output: "c"

    Input:
    letters = ["c", "f", "j"]
    target = "k"
    Output: "c"

  Note:
    1. letters has a length in range [2, 10000].
    2. letters consists of lowercase letters, and contains at least 2 unique letters.
    3. target is a lowercase letter.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n-1
	for l < r {
		m := (l + r) >> 1
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}

	if letters[l] <= target {
		return letters[0]
	}
	return letters[l]
}

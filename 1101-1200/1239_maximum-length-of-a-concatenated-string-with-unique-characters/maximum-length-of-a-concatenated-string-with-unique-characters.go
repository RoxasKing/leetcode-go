package main

/*
  Given an array of strings arr. String s is a concatenation of a sub-sequence of arr which have unique characters.

  Return the maximum possible length of s.

  Example 1:
    Input: arr = ["un","iq","ue"]
    Output: 4
    Explanation: All possible concatenations are "","un","iq","ue","uniq" and "ique".
      Maximum length is 4.

  Example 2:
    Input: arr = ["cha","r","act","ers"]
    Output: 6
    Explanation: Possible solutions are "chaers" and "acters".

  Example 3:
    Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
    Output: 26

  Constraints:
    1. 1 <= arr.length <= 16
    2. 1 <= arr[i].length <= 26
    3. arr[i] contains only lower case English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking

func maxLength(arr []string) int {
	n := len(arr)
	nums := make([]int, n)
	for i := range arr {
		cnt := [26]int{}
		for j := range arr[i] {
			cnt[arr[i][j]-'a']++
			if cnt[arr[i][j]-'a'] > 1 {
				nums[i] = -1
				break
			}
			nums[i] |= 1 << int(arr[i][j]-'a')
		}
	}
	out := 0
	backtrack(arr, nums, n, 0, 0, 0, &out)
	return out
}

func backtrack(arr []string, nums []int, n, i, curVal, curLen int, out *int) {
	if i == n {
		return
	}

	backtrack(arr, nums, n, i+1, curVal, curLen, out)

	if nums[i] != -1 && curVal^nums[i] == curVal+nums[i] {
		curVal += nums[i]
		curLen += len(arr[i])
		*out = Max(*out, curLen)
		backtrack(arr, nums, n, i+1, curVal, curLen, out)
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

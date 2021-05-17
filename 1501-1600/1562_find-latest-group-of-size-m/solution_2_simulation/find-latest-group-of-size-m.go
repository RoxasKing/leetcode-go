package main

/*
  Given an array arr that represents a permutation of numbers from 1 to n. You have a binary string of size n that initially has all its bits set to zero.

  At each step i (assuming both the binary string and arr are 1-indexed) from 1 to n, the bit at position arr[i] is set to 1. You are given an integer m and you need to find the latest step at which there exists a group of ones of length m. A group of ones is a contiguous substring of 1s such that it cannot be extended in either direction.

  Return the latest step at which there exists a group of ones of length exactly m. If no such group exists, return -1.

  Example 1:
    Input: arr = [3,5,1,2,4], m = 1
    Output: 4
    Explanation:
      Step 1: "00100", groups: ["1"]
      Step 2: "00101", groups: ["1", "1"]
      Step 3: "10101", groups: ["1", "1", "1"]
      Step 4: "11101", groups: ["111", "1"]
      Step 5: "11111", groups: ["11111"]
      The latest step at which there exists a group of size 1 is step 4.

  Example 2:
    Input: arr = [3,1,5,4,2], m = 2
    Output: -1
    Explanation:
      Step 1: "00100", groups: ["1"]
      Step 2: "10100", groups: ["1", "1"]
      Step 3: "10101", groups: ["1", "1", "1"]
      Step 4: "10111", groups: ["1", "111"]
      Step 5: "11111", groups: ["11111"]
      No group of size 2 exists during any step.

  Example 3:
    Input: arr = [1], m = 1
    Output: 1

  Example 4:
    Input: arr = [2,1], m = 2
    Output: 2

  Constraints:
    1. n == arr.length
    2. 1 <= n <= 10^5
    3. 1 <= arr[i] <= n
    4. All integers in arr are distinct.
    5. 1 <= m <= arr.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-latest-group-of-size-m
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Simulation

func findLatestStep(arr []int, m int) int {
	n := len(arr)
	gCnt := make([]int, n+1)
	gLen := make([]int, n)
	out := -1

	update := func(l, r int) {
		gCnt[gLen[l]]--
		gCnt[gLen[r]]--
		newLen := gLen[l] + gLen[r]
		gCnt[newLen]++
		gLen[l] = newLen
		gLen[r] = newLen
	}

	for j, i := range arr {
		i--
		j++
		gCnt[1]++
		gLen[i] = 1
		if i-1 >= 0 && gLen[i-1] > 0 {
			l := i - 1 - gLen[i-1] + 1
			r := i - gLen[i] + 1
			update(l, r)
		}
		if i+1 < n && gLen[i+1] > 0 {
			l := i - gLen[i] + 1
			r := i + 1 + gLen[i+1] - 1
			update(l, r)
		}
		if gCnt[m] > 0 {
			out = j
		}
	}

	return out
}

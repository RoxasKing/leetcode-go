package main

import "sort"

/*
  You have a number of envelopes with widths and heights given as a pair of integers (w, h). One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.

  What is the maximum number of envelopes can you Russian doll? (put one inside other)

  Note:
  Rotation is not allowed.

  Example:
    Input: [[5,4],[6,4],[6,7],[2,3]]
    Output: 3
    Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/russian-doll-envelopes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func maxEnvelopes(envelopes [][]int) int {
	n := len(envelopes)
	if n == 0 {
		return 0
	}

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][1] != envelopes[j][1] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] > envelopes[j][0]
	})

	// e.g.          [[12 2] [3 4] [13 3] [4 5] [5 6] [14 4] [15 5] [5 5]]
	// after sort => [[12 2] [13 3] [14 4] [3 4] [15 5] [5 5] [4 5] [5 6]]
	// every heights can only selete one, and the width must be strictly incremented,
	// so we can get arrays like that:
	//               [[12 2] [13 3] [14 4] [15 5]]
	//            or [[3 4] [5 5]]
	//            or [[3 4] [4 5] [5 6]]
	// select the longest one

	arr := make([]int, 0, n)
	for _, envelope := range envelopes {
		num := envelope[0]
		// find the first elem in the array that is not less than the num, return index
		i := sort.Search(len(arr), func(i int) bool { return arr[i] >= num })
		if i < len(arr) {
			arr[i] = num
		} else {
			arr = append(arr, num)
		}
	}
	return len(arr)
}

package main

import "sort"

/*
  You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.

  One envelope can fit into another if and only if both the width and height of one envelope are greater than the other envelope's width and height.

  Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).

  Note: You cannot rotate an envelope.

  Example 1:
    Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
    Output: 3
    Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).

  Example 2:
    Input: envelopes = [[1,1],[1,1],[1,1]]
    Output: 1

  Constraints:
    1. 1 <= envelopes.length <= 5000
    2. envelopes[i].length == 2
    3. 1 <= wi, hi <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/russian-doll-envelopes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Binary Search
func maxEnvelopes(envelopes [][]int) int {
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

	arr := []int{}
	for _, envelope := range envelopes {
		i := sort.Search(len(arr), func(i int) bool { return arr[i] >= envelope[0] })
		if i < len(arr) {
			arr[i] = envelope[0]
			continue
		}
		arr = append(arr, envelope[0])
	}
	return len(arr)
}

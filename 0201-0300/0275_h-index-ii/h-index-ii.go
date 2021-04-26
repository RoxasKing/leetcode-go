package main

/*
  Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper and citations is sorted in an ascending order, return compute the researcher's h-index.

  According to the definition of h-index on Wikipedia: A scientist has an index h if h of their n papers have at least h citations each, and the other n − h papers have no more than h citations each.

  If there are several possible values for h, the maximum one is taken as the h-index.

  Example 1:
    Input: citations = [0,1,3,5,6]
    Output: 3
    Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had received 0, 1, 3, 5, 6 citations respectively.
      Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.

  Example 2:
    Input: citations = [1,2,100]
    Output: 2

  Constraints:
    1. n == citations.length
    2. 1 <= n <= 10^5
    3. 0 <= citations[i] <= 1000
    4. citations is sorted in ascending order.

  Follow up: Could you solve it in logarithmic time complexity?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/h-index-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n
	for l < r {
		m := (l + r) >> 1
		if citations[m] < n-m {
			l = m + 1
		} else {
			r = m
		}
	}
	return n - l
}

package main

import (
	"math"
	"strconv"
)

/*
  Given an integer n represented as a string, return the smallest good base of n.

  We call k >= 2 a good base of n, if all digits of n base k are 1's.

  Example 1:
    Input: n = "13"
    Output: "3"
    Explanation: 13 base 3 is 111.

  Example 2:
    Input: n = "4681"
    Output: "8"
    Explanation: 4681 base 8 is 11111.

  Example 3:
    Input: n = "1000000000000000000"
    Output: "999999999999999999"
    Explanation: 1000000000000000000 base 999999999999999999 is 11.

  Constraints:
    1. n is an integer in the range [3, 1018].
    2. n does not contain any leading zeros.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/smallest-good-base
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math + Binary Search
func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	for i := int(math.Log2(float64(num))) + 1; i >= 2; i-- {
		l, r := 2, int(math.Pow(float64(num), 1/float64(i)))
		for l <= r {
			k := (l + r) >> 1
			sum := 1
			for j := 0; j < i; j++ {
				sum = sum*k + 1
			}
			if sum < num {
				l = k + 1
			} else if sum > num {
				r = k - 1
			} else {
				return strconv.Itoa(k)
			}
		}
	}
	return strconv.Itoa(num - 1)
}

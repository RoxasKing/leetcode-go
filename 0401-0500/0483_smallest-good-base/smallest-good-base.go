package main

import (
	"math"
	"math/bits"
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
    1. n is an integer in the range [3, 10^18].
    2. n does not contain any leading zeros.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/smallest-good-base
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func smallestGoodBase(n string) string {
	nVal, _ := strconv.Atoi(n)
	for m := bits.Len(uint(nVal)) - 1; m > 1; m-- {
		k := int(math.Pow(float64(nVal), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == nVal {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(nVal - 1)
}

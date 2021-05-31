package main

/*
  You are given a very large integer n, represented as a string,​​​​​​ and an integer digit x. The digits in n and the digit x are in the inclusive range [1, 9], and n may represent a negative number.

  You want to maximize n's numerical value by inserting x anywhere in the decimal representation of n​​​​​​. You cannot insert x to the left of the negative sign.

    1. For example, if n = 73 and x = 6, it would be best to insert it between 7 and 3, making n = 763.
    2. If n = -55 and x = 2, it would be best to insert it before the first 5, making n = -255.

  Return a string representing the maximum value of n​​​​​​ after the insertion.

  Example 1:
    Input: n = "99", x = 9
    Output: "999"
    Explanation: The result is the same regardless of where you insert 9.

  Example 2:
    Input: n = "-13", x = 2
    Output: "-123"
    Explanation: You can make n one of {-213, -123, -132}, and the largest of those three is -123.

  Constraints:
    1. 1 <= n.length <= 10^5
    2. 1 <= x <= 9
    3. The digits in n​​​ are in the range [1, 9].
    4. n is a valid representation of an integer.
    5. In the case of a negative n,​​​​​​ it will begin with '-'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-value-after-insertion
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm

func maxValue(n string, x int) string {
	flg := 1
	i := 0
	if n[0] == '-' {
		flg = -flg
		i++
	}
	for ; i < len(n); i++ {
		if flg*x > flg*int(n[i]-'0') {
			break
		}
	}
	return n[:i] + string('0'+byte(x)) + n[i:]
}

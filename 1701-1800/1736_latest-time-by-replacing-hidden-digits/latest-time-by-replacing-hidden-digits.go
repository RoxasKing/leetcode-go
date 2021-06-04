package main

/*
  You are given a string time in the form of hh:mm, where some of the digits in the string are hidden (represented by ?).

  The valid times are those inclusively between 00:00 and 23:59.

  Return the latest valid time you can get from time by replacing the hidden digits.

  Example 1:
    Input: time = "2?:?0"
    Output: "23:50"
    Explanation: The latest hour beginning with the digit '2' is 23 and the latest minute ending with the digit '0' is 50.

  Example 2:
    Input: time = "0?:3?"
    Output: "09:39"

  Example 3:
    Input: time = "1?:22"
    Output: "19:22"

  Constraints:
    1. time is in the format hh:mm.
    2. It is guaranteed that you can produce a valid time from the given string.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximumTime(time string) string {
	cs := []byte(time)
	if cs[0] == '?' {
		if '4' <= cs[1] && cs[1] <= '9' {
			cs[0] = '1'
		} else {
			cs[0] = '2'
		}
	}
	if cs[1] == '?' {
		if cs[0] == '2' {
			cs[1] = '3'
		} else {
			cs[1] = '9'
		}
	}
	if cs[3] == '?' {
		cs[3] = '5'
	}
	if cs[4] == '?' {
		cs[4] = '9'
	}
	return string(cs)
}

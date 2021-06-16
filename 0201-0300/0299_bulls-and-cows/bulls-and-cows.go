package main

import "strconv"

/*
  You are playing the Bulls and Cows game with your friend.

  You write down a secret number and ask your friend to guess what the number is. When your friend makes a guess, you provide a hint with the following info:

    The number of "bulls", which are digits in the guess that are in the correct position.
    The number of "cows", which are digits in the guess that are in your secret number but are located in the wrong position. Specifically, the non-bull digits in the guess that could be rearranged such that they become bulls.

  Given the secret number secret and your friend's guess guess, return the hint for your friend's guess.

  The hint should be formatted as "xAyB", where x is the number of bulls and y is the number of cows. Note that both secret and guess may contain duplicate digits.

  Example 1:
    Input: secret = "1807", guess = "7810"
    Output: "1A3B"
    Explanation: Bulls are connected with a '|' and cows are underlined:
      "1807"
        |
      "7810"

  Example 2:
    Input: secret = "1123", guess = "0111"
    Output: "1A1B"
    Explanation: Bulls are connected with a '|' and cows are underlined:
      "1123"        "1123"
        |      or     |
      "0111"        "0111"
      Note that only one of the two unmatched 1s is counted as a cow since the non-bull digits can only be rearranged to allow one 1 to be a bull.

  Example 3:
    Input: secret = "1", guess = "0"
    Output: "0A0B"

  Example 4:
    Input: secret = "1", guess = "1"
    Output: "1A0B"

  Constraints:
    1. 1 <= secret.length, guess.length <= 1000
    2. secret.length == guess.length
    3. secret and guess consist of digits only.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/bulls-and-cows
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func getHint(secret string, guess string) string {
	a, b := 0, 0
	h := [10]int{}
	for i := range secret {
		h[secret[i]-'0']++
	}
	for i := range guess {
		if guess[i] == secret[i] {
			a++
			h[secret[i]-'0']--
		}
	}
	for i := range guess {
		if guess[i] == secret[i] {
			continue
		}
		if h[guess[i]-'0'] > 0 {
			b++
			h[guess[i]-'0']--
		}
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

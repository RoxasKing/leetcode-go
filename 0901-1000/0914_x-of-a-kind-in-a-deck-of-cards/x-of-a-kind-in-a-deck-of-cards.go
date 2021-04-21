package main

/*
  In a deck of cards, each card has an integer written on it.

  Return true if and only if you can choose X >= 2 such that it is possible to split the entire deck into 1 or more groups of cards, where:
    1. Each group has exactly X cards.
    2. All the cards in each group have the same integer.


  Example 1:
    Input: deck = [1,2,3,4,4,3,2,1]
    Output: true
    Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].

  Example 2:
    Input: deck = [1,1,1,2,2,2,3,3]
    Output: false
    Explanation: No possible partition.

  Example 3:
    Input: deck = [1]
    Output: false
    Explanation: No possible partition.

  Example 4:
    Input: deck = [1,1]
    Output: true
    Explanation: Possible partition [1,1].

  Example 5:
    Input: deck = [1,1,2,2,2,2]
    Output: true
    Explanation: Possible partition [1,1],[2,2],[2,2].

  Constraints:
    1. 1 <= deck.length <= 10^4
    2. 0 <= deck[i] < 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int)
	for _, num := range deck {
		cnt[num]++
	}
	prev := -1
	for _, v := range cnt {
		if prev != -1 {
			prev = Gcd(prev, v)
		} else {
			prev = v
		}
	}
	return prev >= 2
}

func Gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

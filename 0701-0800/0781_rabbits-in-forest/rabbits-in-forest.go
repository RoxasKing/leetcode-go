package main

/*
  In a forest, each rabbit has some color. Some subset of rabbits (possibly all of them) tell you how many other rabbits have the same color as them. Those answers are placed in an array.

  Return the minimum number of rabbits that could be in the forest.

  Examples:
    Input: answers = [1, 1, 2]
    Output: 5
    Explanation:
      The two rabbits that answered "1" could both be the same color, say red.
      The rabbit than answered "2" can't be red or the answers would be inconsistent.
      Say the rabbit that answered "2" was blue.
      Then there should be 2 other blue rabbits in the forest that didn't answer into the array.
      The smallest possible number of rabbits in the forest is therefore 5: 3 that answered plus 2 that didn't.

    Input: answers = [10, 10, 10]
    Output: 11

    Input: answers = []
    Output: 0

  Note:
    1. answers will have length at most 1000.
    2. Each answers[i] will be an integer in the range [0, 999].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rabbits-in-forest
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math + Hash
func numRabbits(answers []int) int {
	dict := make(map[int]int)
	for _, a := range answers {
		dict[a]++
	}
	out := 0
	for k, v := range dict {
		if k+1 >= v {
			out += k + 1
		} else {
			t := v / (k + 1)
			out += t * (k + 1)
			if v%(k+1) > 0 {
				out += k + 1
			}
		}
	}
	return out
}

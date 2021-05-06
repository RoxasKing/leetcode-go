package main

/*
  On a social network consisting of m users and some friendships between users, two users can communicate with each other if they know a common language.

  You are given an integer n, an array languages, and an array friendships where:
    1. There are n languages numbered 1 through n,
    2. languages[i] is the set of languages the i​​​​​​th​​​​ user knows, and
    3. friendships[i] = [u​​​​​​i​​​, v​​​​​​i] denotes a friendship between the users u​​​​​​​​​​​i​​​​​ and vi.

  You can choose one language and teach it to some users so that all friends can communicate with each other. Return the minimum number of users you need to teach.

  Note that friendships are not transitive, meaning if x is a friend of y and y is a friend of z, this doesn't guarantee that x is a friend of z.


  Example 1:
    Input: n = 2, languages = [[1],[2],[1,2]], friendships = [[1,2],[1,3],[2,3]]
    Output: 1
    Explanation: You can either teach user 1 the second language or user 2 the first language.

  Example 2:
    Input: n = 3, languages = [[2],[1,3],[1,2],[3]], friendships = [[1,4],[1,2],[3,4],[2,3]]
    Output: 2
    Explanation: Teach the third language to users 1 and 3, yielding two users to teach.

  Constraints:
    1. 2 <= n <= 500
    2. languages.length == m
    3. 1 <= m <= 500
    4. 1 <= languages[i].length <= n
    5. 1 <= languages[i][j] <= n
    6. 1 <= u​​​​​​i < v​​​​​​i <= languages.length
    7. 1 <= friendships.length <= 500
    8. All tuples (u​​​​​i, v​​​​​​i) are unique
    9. languages[i] contains only unique values

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-number-of-people-to-teach
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	m := len(languages)
	cnt := make([]int, n+1)
	mark := make(map[int]bool)

	for _, f := range friendships {
		u, v := f[0], f[1]
		can := make([]bool, n+1)
		skip := false
		for _, lan := range languages[u-1] {
			can[lan] = true
		}
		for _, lan := range languages[v-1] {
			if can[lan] {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		if !mark[u] {
			for _, lan := range languages[u-1] {
				cnt[lan]++
			}
			mark[u] = true
		}
		if !mark[v] {
			for _, lan := range languages[v-1] {
				cnt[lan]++
			}
			mark[v] = true
		}
	}

	out := m
	for i := 1; i <= n; i++ {
		out = Min(out, len(mark)-cnt[i])
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

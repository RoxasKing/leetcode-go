package main

/*
  You are an ant tasked with adding n new rooms numbered 0 to n-1 to your colony. You are given the expansion plan as a 0-indexed integer array of length n, prevRoom, where prevRoom[i] indicates that you must build room prevRoom[i] before building room i, and these two rooms must be connected directly. Room 0 is already built, so prevRoom[0] = -1. The expansion plan is given such that once all the rooms are built, every room will be reachable from room 0.

  You can only build one room at a time, and you can travel freely between rooms you have already built only if they are connected. You can choose to build any room as long as its previous room is already built.

  Return the number of different orders you can build all the rooms in. Since the answer may be large, return it modulo 10^9 + 7.

  Example 1:
    Input: prevRoom = [-1,0,1]
    Output: 1
    Explanation: There is only one way to build the additional rooms: 0 → 1 → 2

  Example 2:
    Input: prevRoom = [-1,0,0,1,2]
    Output: 6
    Explanation:
      The 6 ways are:
      0 → 1 → 3 → 2 → 4
      0 → 2 → 4 → 1 → 3
      0 → 1 → 2 → 3 → 4
      0 → 1 → 2 → 4 → 3
      0 → 2 → 1 → 3 → 4
      0 → 2 → 1 → 4 → 3

  Constraints:
    1. n == prevRoom.length
    2. 2 <= n <= 10^5
    3. prevRoom[0] == -1
    4. 0 <= prevRoom[i] < n for all 1 <= i < n
    5. Every room is reachable from room 0 once all the rooms are built.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-ways-to-build-rooms-in-an-ant-colony
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Topological sorting
// Math

func waysToBuildRooms(prevRoom []int) int {
	n := len(prevRoom)
	fac := make([]int, n+1)
	ifac := make([]int, n+1)
	fac[0], fac[1] = 1, 1
	ifac[0], ifac[1] = 1, 1
	for i := 2; i <= n; i++ {
		fac[i] = fac[i-1] * i % mod
		ifac[i] = quickMul(fac[i], mod-2)
	}

	edges := make([][]int, n)
	indeg := make([]int, n)
	for i := 1; i < n; i++ {
		edges[prevRoom[i]] = append(edges[prevRoom[i]], i)
		indeg[prevRoom[i]]++
	}

	lens := make([]int, n)
	for i := range lens {
		lens[i] = 1
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	f := make([]int, n)

	for ; len(q) > 0; q = q[1:] {
		u := q[0]
		p := prevRoom[u]
		if 0 <= p {
			lens[p] += lens[u]
			indeg[p]--
			if indeg[p] == 0 {
				q = append(q, p)
			}
		}
		cnt := fac[lens[u]-1]
		for _, v := range edges[u] {
			cnt = cnt * ifac[lens[v]] % mod
			cnt = cnt * f[v] % mod
		}
		f[u] = cnt
	}

	return f[0]
}

func quickMul(x, n int) int {
	out := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			out = out * x % mod
		}
		x = x * x % mod
	}
	return out
}

var mod = int(1e9 + 7)

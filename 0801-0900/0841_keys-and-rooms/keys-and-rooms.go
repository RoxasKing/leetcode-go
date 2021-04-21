package main

/*
  有 N 个房间，开始时你位于 0 号房间。每个房间有不同的号码：0，1，2，...，N-1，并且房间里可能有一些钥匙能使你进入下一个房间。

  在形式上，对于每个房间 i 都有一个钥匙列表 rooms[i]，每个钥匙 rooms[i][j] 由 [0,1，...，N-1] 中的一个整数表示，其中 N = rooms.length。 钥匙 rooms[i][j] = v 可以打开编号为 v 的房间。

  最初，除 0 号房间外的其余所有房间都被锁住。

  你可以自由地在房间之间来回走动。

  如果能进入每个房间返回 true，否则返回 false。

  提示：
    1 <= rooms.length <= 1000
    0 <= rooms[i].length <= 1000
    所有房间中的钥匙数量总计不超过 3000。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/keys-and-rooms
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursion
func canVisitAllRooms(rooms [][]int) bool {
	mark := make([]bool, len(rooms))
	var dfs func(int)
	dfs = func(index int) {
		if mark[index] {
			return
		}
		mark[index] = true
		for _, key := range rooms[index] {
			dfs(key)
		}
	}
	dfs(0)
	for i := range mark {
		if !mark[i] {
			return false
		}
	}
	return true
}

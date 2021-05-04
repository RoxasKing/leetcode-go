package main

import (
	"sort"
)

/*
  There is a hotel with n rooms. The rooms are represented by a 2D integer array rooms where rooms[i] = [roomIdi, sizei] denotes that there is a room with room number roomIdi and size equal to sizei. Each roomIdi is guaranteed to be unique.

  You are also given k queries in a 2D array queries where queries[j] = [preferredj, minSizej]. The answer to the jth query is the room number id of a room such that:
    1. The room has a size of at least minSizej, and
    2. abs(id - preferredj) is minimized, where abs(x) is the absolute value of x.
  If there is a tie in the absolute difference, then use the room with the smallest such id. If there is no such room, the answer is -1.

  Return an array answer of length k where answer[j] contains the answer to the jth query.

  Example 1:
    Input: rooms = [[2,2],[1,2],[3,2]], queries = [[3,1],[3,3],[5,2]]
    Output: [3,-1,3]
    Explanation: The answers to the queries are as follows:
      Query = [3,1]: Room number 3 is the closest as abs(3 - 3) = 0, and its size of 2 is at least 1. The answer is 3.
      Query = [3,3]: There are no rooms with a size of at least 3, so the answer is -1.
      Query = [5,2]: Room number 3 is the closest as abs(3 - 5) = 2, and its size of 2 is at least 2. The answer is 3.

  Example 2:
    Input: rooms = [[1,4],[2,3],[3,5],[4,1],[5,2]], queries = [[2,3],[2,4],[2,5]]
    Output: [2,1,3]
    Explanation: The answers to the queries are as follows:
      Query = [2,3]: Room number 2 is the closest as abs(2 - 2) = 0, and its size of 3 is at least 3. The answer is 2.
      Query = [2,4]: Room numbers 1 and 3 both have sizes of at least 4. The answer is 1 since it is smaller.
      Query = [2,5]: Room number 3 is the only room with a size of at least 5. The answer is 3.

  Constraints:
    1. n == rooms.length
    2. 1 <= n <= 10^5
    3. k == queries.length
    4. 1 <= k <= 10^4
    5. 1 <= roomIdi, preferredj <= 10^7
    6. 1 <= sizei, minSizej <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/closest-room
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

func closestRoom(rooms [][]int, queries [][]int) []int {
	n := len(queries)

	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool { return queries[idxs[i]][1] < queries[idxs[j]][1] })
	sort.Slice(rooms, func(i, j int) bool { return rooms[i][1] < rooms[j][1] })

	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = -1
	}

	for _, idx := range idxs {
		q := queries[idx]
		i := sort.Search(len(rooms), func(i int) bool { return rooms[i][1] >= q[1] })
		rooms = rooms[i:]
		for _, room := range rooms {
			if out[idx] == -1 || Abs(out[idx]-q[0]) > Abs(room[0]-q[0]) ||
				Abs(out[idx]-q[0]) == Abs(room[0]-q[0]) && room[0] < out[idx] {
				out[idx] = room[0]
			}
			if Abs(out[idx]-q[0]) == 0 {
				break
			}
		}
	}

	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

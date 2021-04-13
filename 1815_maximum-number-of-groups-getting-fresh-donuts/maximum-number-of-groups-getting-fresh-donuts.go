package main

/*
  There is a donuts shop that bakes donuts in batches of batchSize. They have a rule where they must serve all of the donuts of a batch before serving any donuts of the next batch. You are given an integer batchSize and an integer array groups, where groups[i] denotes that there is a group of groups[i] customers that will visit the shop. Each customer will get exactly one donut.

  When a group visits the shop, all customers of the group must be served before serving any of the following groups. A group will be happy if they all get fresh donuts. That is, the first customer of the group does not receive a donut that was left over from the previous group.

  You can freely rearrange the ordering of the groups. Return the maximum possible number of happy groups after rearranging the groups.

  Example 1:
    Input: batchSize = 3, groups = [1,2,3,4,5,6]
    Output: 4
    Explanation: You can arrange the groups as [6,2,4,5,1,3]. Then the 1st, 2nd, 4th, and 6th groups will be happy.

  Example 2:
    Input: batchSize = 4, groups = [1,3,2,5,2,2,1,6]
    Output: 4

  Constraints:
    1. 1 <= batchSize <= 9
    2. 1 <= groups.length <= 30
    3. 1 <= groups[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-number-of-groups-getting-fresh-donuts
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func maxHappyGroups(batchSize int, groups []int) int {
	freq0 := make([]int, 9)
	freq := make([]int, 9)
	w := make([]int, 9)

	for _, g := range groups {
		freq0[g%batchSize]++
	}

	msum := 1
	for i := 1; i < batchSize; i++ {
		msum *= (freq0[i] + 1)
	}

	w[1] = 1
	for i := 2; i < batchSize; i++ {
		w[i] = w[i-1] * (freq0[i-1] + 1)
	}

	f := make([]int, msum)
	for fmask := 1; fmask < msum; fmask++ {
		last := 0
		for i := 1; i < batchSize; i++ {
			freq[i] = (fmask / w[i]) % (freq0[i] + 1)
			last = (last + (freq0[i]-freq[i])*i) % batchSize
		}
		for c := 1; c < batchSize; c++ {
			if freq[c] == 0 {
				continue
			}
			cnt := f[fmask-w[c]]
			if last == 0 {
				cnt++
			}
			f[fmask] = Max(f[fmask], cnt)
		}
	}
	return f[msum-1] + freq0[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

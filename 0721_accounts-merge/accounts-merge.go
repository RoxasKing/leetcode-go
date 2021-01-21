package main

import "sort"

/*
  Given a list accounts, each element accounts[i] is a list of strings, where the first element accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.

  Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there is some email that is common to both accounts. Note that even if two accounts have the same name, they may belong to different people as people could have the same name. A person can have any number of accounts initially, but all of their accounts definitely have the same name.

  After merging the accounts, return the accounts in the following format: the first element of each account is the name, and the rest of the elements are emails in sorted order. The accounts themselves can be returned in any order.

  Example 1:
  Input:
    accounts = [["John", "johnsmith@mail.com", "john00@mail.com"], ["John", "johnnybravo@mail.com"], ["John", "johnsmith@mail.com", "john_newyork@mail.com"], ["Mary", "mary@mail.com"]]
  Output:
    [["John", 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com'],  ["John", "johnnybravo@mail.com"], ["Mary", "mary@mail.com"]]
  Explanation:
    The first and third John's are the same person as they have the common email "johnsmith@mail.com".
    The second John and Mary are different people as none of their email addresses are used by other accounts.
    We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'], ['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.

  Note:
    The length of accounts will be in the range [1, 1000].
    The length of accounts[i] will be in the range [1, 10].
    The length of accounts[i][j] will be in the range [1, 30].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/accounts-merge
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Union-Find
func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)

	// union all related accounts
	uf := newUnionFind(n)

	// record all emails account index
	emails := make(map[string]int)

	for i, account := range accounts {
		for _, email := range account[1:] {
			if idx, ok := emails[email]; ok {
				uf.union(idx, i)
			} else {
				emails[email] = i
			}
		}
	}

	accountsMap := make(map[int][]string)

	for i := range accounts {
		idx := uf.find(i)
		if _, ok := accountsMap[idx]; ok {
			continue
		}
		name := accounts[i][0]
		accountsMap[idx] = append(accountsMap[idx], name)
	}

	for email, i := range emails {
		idx := uf.find(i)
		accountsMap[idx] = append(accountsMap[idx], email)
	}

	out := make([][]string, 0, len(accountsMap))
	for _, account := range accountsMap {
		sort.Strings(account)
		out = append(out, account)
	}

	return out
}

type unionFind struct {
	ancestor []int
	isEnd    []bool
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := range ancestor {
		ancestor[i] = i
	}
	isEnd := make([]bool, n)
	return unionFind{ancestor: ancestor, isEnd: isEnd}
}

func (uf unionFind) find(x int) int {
	if uf.isEnd[uf.ancestor[x]] {
		return uf.ancestor[x]
	}
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.find(uf.ancestor[x])
		uf.isEnd[x] = false
		uf.isEnd[uf.ancestor[x]] = true
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(x, y int) {
	ancestorX := uf.find(x)
	ancestorY := uf.find(y)
	uf.ancestor[ancestorX] = ancestorY
	uf.isEnd[ancestorX] = false
	uf.isEnd[ancestorY] = true
}

package main

import (
	"strings"
)

/*
  One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as #.
       _9_
      /   \
     3     2
    / \   / \
   4   1  #  6
  / \ / \   / \
  # # # #   # #
  For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", where # represents a null node.

  Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.

  Each comma separated value in the string must be either an integer or a character '#' representing null pointer.

  You may assume that the input format is always valid, for example it could never contain two consecutive commas such as "1,,3".

  Example 1:
    Input: "9,3,4,#,#,1,#,#,2,#,6,#,#"
    Output: true

  Example 2:
    Input: "1,#"
    Output: false

  Example 3:
    Input: "9,#,#,1"
    Output: false

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func isValidSerialization(preorder string) bool {
	strs := strings.Split(preorder, ",")
	i := 0
	dfs(strs, &i)
	return i == len(strs)-1
}

func dfs(strs []string, i *int) bool {
	if strs[*i] == "#" {
		return true
	}
	*i++
	if *i == len(strs) {
		return false
	}
	ok := dfs(strs, i)
	if !ok {
		return false
	}
	*i++
	if *i == len(strs) {
		return false
	}
	ok = dfs(strs, i)
	return ok
}

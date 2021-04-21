package main

import (
	"strconv"
	"strings"
)

/*
  序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

  请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

  提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

  说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := []*TreeNode{root}
	var strs []string
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			strs = append(strs, "#")
		} else {
			strs = append(strs, strconv.Itoa(node.Val))
			q = append(q, node.Left, node.Right)
		}
	}
	return strings.Join(strs, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	if strs[0] == "#" {
		return nil
	}
	rootVal, _ := strconv.Atoi(strs[0])
	strs = strs[1:]
	root := &TreeNode{Val: rootVal}
	q := []*TreeNode{root}
	for len(q) != 0 && len(strs) != 0 {
		node := q[0]
		q = q[1:]
		if strs[0] != "#" {
			lVal, _ := strconv.Atoi(strs[0])
			lNode := &TreeNode{Val: lVal}
			node.Left = lNode
			q = append(q, lNode)
		}
		if strs[1] != "#" {
			rVal, _ := strconv.Atoi(strs[1])
			rNode := &TreeNode{Val: rVal}
			node.Right = rNode
			q = append(q, rNode)
		}
		strs = strs[2:]
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

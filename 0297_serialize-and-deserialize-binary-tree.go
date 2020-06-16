package leetcode

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

func NewCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "$,"
	}
	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + c.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	strs := strings.Split(data, ",")
	var index int
	var build func() *TreeNode
	build = func() *TreeNode {
		if index == len(strs)-1 {
			return nil
		}
		if strs[index] == "$" {
			index++
			return nil
		}
		rootVal, _ := strconv.Atoi(strs[index])
		index++
		return &TreeNode{
			Val:   rootVal,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}

func (c *Codec) serialize2(root *TreeNode) string {
	stack := []*TreeNode{root}
	var strs []string
	for len(stack) != 0 {
		root, stack = stack[0], stack[1:]
		if root == nil {
			strs = append(strs, "$")
		} else {
			strs = append(strs, strconv.Itoa(root.Val))
			stack = append(stack, root.Left, root.Right)
		}
	}
	return strings.Join(strs, ",")
}

func (c *Codec) deserialize2(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var index int
	strs := strings.Split(data, ",")
	rootVal, _ := strconv.Atoi(strs[index])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if strs[index+1] != "$" {
			val, _ := strconv.Atoi(strs[index+1])
			l := &TreeNode{Val: val}
			node.Left = l
			queue = append(queue, l)
		}
		if strs[index+2] != "$" {
			val, _ := strconv.Atoi(strs[index+2])
			r := &TreeNode{Val: val}
			node.Right = r
			queue = append(queue, r)
		}
		index += 2
	}
	return root
}

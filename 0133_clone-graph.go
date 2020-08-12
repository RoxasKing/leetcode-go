package leetcode

/*
  给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。
  图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

  class Node {
      public int val;
      public List<Node> neighbors;
  }

  测试用例格式：
  简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
  邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
  给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/clone-graph
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// Hash + BFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	mark := map[int]*Node{node.Val: {Val: node.Val}}
	srcQ := []*Node{node}
	dstQ := []*Node{mark[node.Val]}
	for len(srcQ) != 0 {
		srcN, dstN := srcQ[0], dstQ[0]
		srcQ, dstQ = srcQ[1:], dstQ[1:]
		for _, srcn := range srcN.Neighbors {
			if _, ok := mark[srcn.Val]; !ok {
				mark[srcn.Val] = &Node{Val: srcn.Val}
				srcQ = append(srcQ, srcn)
				dstQ = append(dstQ, mark[srcn.Val])
			}
			dstN.Neighbors = append(dstN.Neighbors, mark[srcn.Val])
		}
	}
	return mark[node.Val]
}

// Hash + BFS
func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := []*Node{node}
	mark := map[*Node]*Node{node: {Val: node.Val}}
	for len(queue) != 0 {
		N := queue[0]
		queue = queue[1:]
		for _, n := range N.Neighbors {
			if _, ok := mark[n]; !ok {
				mark[n] = &Node{Val: n.Val}
				queue = append(queue, n)
			}
			mark[N].Neighbors = append(mark[N].Neighbors, mark[n])
		}
	}
	return mark[node]
}

type Node struct {
	Val       int
	Neighbors []*Node
}

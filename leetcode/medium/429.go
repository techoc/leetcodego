package medium

/* Definition for a Node.
* type Node struct {
	*     Val int
	*     Children []*Node
	* }
*/

//给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
//输入：root = [1,null,3,2,4,null,5,6]
//输出：[[1],[3,2,4],[5,6]]

type node struct {
	Val      int
	Children []*node
}

func levelOrder(root *node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var queue []*node
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		var size = len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		res = append(res, level)
	}
	return res
}

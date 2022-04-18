package easy

//给定一个 n叉树的根节点root，返回 其节点值的 后序遍历 。
//n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node590 struct {
	Val      int
	Children []*Node590
}

func postorder(root *Node590) (ans []int) {
	var dfs func(*Node590)
	dfs = func(node *Node590) {
		if node == nil {
			return
		}
		for _, ch := range node.Children {
			dfs(ch)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}

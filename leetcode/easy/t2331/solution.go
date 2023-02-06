package t2331

// https://leetcode.cn/problems/evaluate-boolean-binary-tree/
func evaluateTree(root *TreeNode) bool {
	// 判断是否是叶子节点
	if root.Left == nil {
		return root.Val == 1
	}
	// 遍历左右子树
	l, r := evaluateTree(root.Left), evaluateTree(root.Right)

	// 如果值为2
	if root.Val == 2 {
		return l || r
	}
	return l && r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

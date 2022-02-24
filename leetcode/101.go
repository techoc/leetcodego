package leetcode

//给你一个二叉树的根节点 root ， 检查它是否轴对称。

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var checkSame func(left, right *TreeNode) bool
	checkSame = func(left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}
		if left.Val != right.Val {
			return false
		}
		return checkSame(left.Left, right.Right) && checkSame(left.Right, right.Left)
	}
	return checkSame(root.Left, root.Right)
}

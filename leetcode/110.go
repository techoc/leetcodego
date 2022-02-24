package leetcode

import "math"

//给定一个二叉树，判断它是否是高度平衡的二叉树。
//本题中，一棵高度平衡二叉树定义为：
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var getDepth func(root *TreeNode) int
	getDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := getDepth(root.Left)
		right := getDepth(root.Right)
		//返回节点左右子树的最大深度+1
		return max(left, right) + 1
	}

	if math.Abs(float64(getDepth(root.Left)-getDepth(root.Right))) > 1 {
		return false
	}
	// 左右子树都是平衡二叉树
	return isBalanced(root.Left) && isBalanced(root.Right)
}

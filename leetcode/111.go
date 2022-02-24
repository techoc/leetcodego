package leetcode

//给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量
//说明：叶子节点是指没有子节点的节点。

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil {
		//如果左子树为空，则右子树的最小深度+1
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		//如果右子树为空，则左子树的最小深度+1
		return minDepth(root.Left) + 1
	}
	//如果左右子树都不为空，则最小深度为左右子树的最小值+1
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

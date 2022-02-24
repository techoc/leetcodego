package leetcode

//给你二叉树的根节点root 和一个表示目标和的整数targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。如果存在，返回 true ；否则，返回 false 。
//叶子节点 是指没有子节点的节点。

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(root *TreeNode, targetSum int) bool
	dfs = func(root *TreeNode, targetSum int) bool {
		if root == nil {
			return false
		}
		if root.Val == targetSum && root.Left == nil && root.Right == nil {
			//如果节点等于目标值，并且是叶子节点，则说明存在路径
			return true
		}
		//递归寻找左右子树是否有节点等于目标值
		return dfs(root.Left, targetSum-root.Val) || dfs(root.Right, targetSum-root.Val)
	}
	return dfs(root, targetSum)
}

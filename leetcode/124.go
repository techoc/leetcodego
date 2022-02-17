package leetcode

import "math"

//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
//同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//路径和 是路径中各节点值的总和。
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 如果子树路径和为负则应当置0表示最大路径不包含子树
		left := max(0, maxGain(node.Left))
		right := max(0, maxGain(node.Right))
		// left->root->right 作为路径与已经计算过历史最大值做比较
		maxSum = max(maxSum, left+right+node.Val)
		//选择最大的子树加上当前的节点
		return node.Val + max(left, right)
	}
	maxGain(root)
	return maxSum
}

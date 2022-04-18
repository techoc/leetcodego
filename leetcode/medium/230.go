package medium

import (
	"github.com/techoc/leetcodego/leetcode/hard"
)

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3

func kthSmallest(root *hard.TreeNode, k int) int {
	if root == nil {
		return 0
	}
	//遍历二叉树
	var res []int
	var dfs func(root *hard.TreeNode)
	dfs = func(root *hard.TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	//找到第k个最小的元素
	return res[k-1]
}

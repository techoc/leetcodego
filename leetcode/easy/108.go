package easy

import (
	"github.com/techoc/leetcodego/leetcode/hard"
)

//给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1」的二叉树。

func sortedArrayToBST(nums []int) *hard.TreeNode {
	var buildTree func(nums []int, start, end int) *hard.TreeNode
	buildTree = func(nums []int, start, end int) *hard.TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		root := &hard.TreeNode{Val: nums[mid]}
		root.Left = buildTree(nums, start, mid-1)
		root.Right = buildTree(nums, mid+1, end)
		return root
	}
	return buildTree(nums, 0, len(nums)-1)
}

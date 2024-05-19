package t104

import "github.com/techoc/leetcodego/leetcode/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 记录最大深度
var res int

// 记录遍历到的节点的深度
var depth int

// 104. 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *utils.TreeNode) int {
	// 初始化最大深度和节点深度 避免全局变量污染
	res, depth = 0, 0
	traverse(root)
	return res
}

// 遍历二叉树
func traverse(root *utils.TreeNode) {
	if root == nil {
		return
	}
	// 前序位置 节点深度加一
	depth++
	// 节点为叶子节点时，更新最大深度
	if root.Left == nil && root.Right == nil {
		res = max(res, depth)
	}
	traverse(root.Left)
	traverse(root.Right)
	// 后序位置 节点深度减一
	depth--
}

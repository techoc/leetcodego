package easy

import (
	"github.com/techoc/leetcodego/leetcode/hard"
)

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。

func inorderTraversal(root *hard.TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

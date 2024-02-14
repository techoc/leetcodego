package t145

import "github.com/techoc/leetcodego/leetcode/utils"

// 145.二叉树的后序遍历
// https://leetcode.cn/problems/binary-tree-postorder-traversal
func postorderTraversal(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int

	// 后序遍历
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

package easy

func postorderTraversal(root *TreeNode) []int {
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

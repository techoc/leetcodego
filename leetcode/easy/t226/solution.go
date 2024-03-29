package t226

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	//递归函数的终止条件，节点为空时返回
	if root == nil {
		return nil
	}
	//交换当前节点的左右子节点
	root.Left, root.Right = root.Right, root.Left
	//继续递归地遍历左右子节点
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

package medium

import (
	"github.com/techoc/leetcodego/leetcode/hard"
)

//给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。

func BuildTree(preorder []int, inorder []int) *hard.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//找到父节点
	root := &hard.TreeNode{
		Val: preorder[0],
	}
	//找到父节点在中序遍历中的位置
	index := 0
	for k, v := range inorder {
		if v == root.Val {
			index = k
			break
		}
	}
	//递归构造左子树
	left := BuildTree(preorder[1:index+1], inorder[:index])
	//递归构造右子树
	right := BuildTree(preorder[index+1:], inorder[index+1:])
	//构造整个树
	root.Left = left
	root.Right = right
	return root
}

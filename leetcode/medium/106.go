package medium

import (
	"github.com/techoc/leetcodego/leetcode/hard"
)

//给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。
//输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//输出：[3,9,20,null,null,15,7]

func buildTree(inorder []int, postorder []int) *hard.TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	//找到父节点
	root := &hard.TreeNode{Val: postorder[len(postorder)-1]}
	//找到父节点在中序遍历中的位置
	index := 0
	for k, v := range inorder {
		if v == root.Val {
			index = k
			break
		}
	}
	//找到左节点个数
	leftNum := index
	//构造左子树
	left := buildTree(inorder[:leftNum], postorder[:leftNum])
	//构造右子树
	right := buildTree(inorder[leftNum+1:], postorder[leftNum:len(postorder)-1])
	//构造父节点
	root.Left = left
	root.Right = right
	return root
}

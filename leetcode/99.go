package leetcode

import "math"

//给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。

func recoverTree(root *TreeNode) {
	//中序遍历比较前后访问的节点值，prev 保存上一个访问的节点，当前访问的是 root 节点。
	//每访问一个节点，如果prev.val>=root.val，就找到了一对“错误对”。
	//检查一下它是第一对错误对，还是第二对错误对。
	//遍历结束，就确定了待交换的两个错误点，进行交换。

	prev := &TreeNode{Val: math.MinInt32 - 1}
	var err1, err2 *TreeNode

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)

		if prev.Val >= root.Val && err1 == nil {
			err1 = prev
		}
		if prev.Val >= root.Val && err1 != nil {
			err2 = root
		}
		prev = root
		inorder(root.Right)
	}

	inorder(root)
	err1.Val, err2.Val = err2.Val, err1.Val
}

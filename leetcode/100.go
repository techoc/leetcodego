package leetcode

//给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var inorder func(p *TreeNode, q *TreeNode) bool
	inorder = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return inorder(p.Left, q.Left) && inorder(p.Right, q.Right)
	}
	return inorder(p, q)
}

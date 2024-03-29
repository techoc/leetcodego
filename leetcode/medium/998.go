package medium

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	//遍历树
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			//如果目标值大于根节点的值
			//则直接构造新树返回
			if parent == nil {
				return &TreeNode{val, root, nil}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{Val: val}
	return root
}

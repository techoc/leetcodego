package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ListToTree(arr []any) *TreeNode {
	return ListToTreeDFS(arr, 0)
}

func ListToTreeDFS(arr []any, i int) *TreeNode {
	if i < 0 || i >= len(arr) || arr[i] == nil {
		return nil
	}
	root := &TreeNode{
		Val: arr[i].(int),
	}
	root.Left = ListToTreeDFS(arr, 2*i+1)
	root.Right = ListToTreeDFS(arr, 2*i+2)
	return root
}

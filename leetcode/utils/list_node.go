package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func listToTree(arr []any) *TreeNode {
	return listToTreeDFS(arr, 0)
}

func listToTreeDFS(arr []any, i int) *TreeNode {
	if i < 0 || i >= len(arr) || arr[i] == nil {
		return nil
	}
	root := &TreeNode{
		Val: arr[i].(int),
	}
	root.Left = listToTreeDFS(arr, 2*i+1)
	root.Right = listToTreeDFS(arr, 2*i+2)
	return root
}

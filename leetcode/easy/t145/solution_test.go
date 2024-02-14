package t145

import "testing"

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: listToTree([]any{1, nil, 2, nil, nil, 3, nil}),
			want: []int{3, 2, 1},
		},
		{
			root: listToTree([]any{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			root: listToTree([]any{}),
			want: []int{},
		},
	}
	for _, tt := range tests {
		got := postorderTraversal(tt.root)
		if !arrayEqual(got, tt.want) {
			t.Errorf("postorderTraversal(%v) = %v, want %v", tt.root, got, tt.want)
		}
	}
}

func arrayEqual(got []int, want []int) bool {
	if len(got) != len(want) {
		return false
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false
		}
	}
	return true
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

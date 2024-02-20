package t145

import (
	"github.com/techoc/leetcodego/leetcode/utils"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		root *utils.TreeNode
		want []int
	}{
		{
			root: utils.ListToTree([]any{1, nil, 2, nil, nil, 3, nil}),
			want: []int{3, 2, 1},
		},
		{
			root: utils.ListToTree([]any{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			root: utils.ListToTree([]any{}),
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

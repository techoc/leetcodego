package t543

import (
	"github.com/techoc/leetcodego/leetcode/utils"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		root *utils.TreeNode
		want int
	}{
		{
			name: "Case 1",
			root: utils.ListToTree([]any{1, 2, 3, 4, 5}),
			want: 3,
		},
		{
			name: "Case 2",
			root: utils.ListToTree([]any{1, 2}),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

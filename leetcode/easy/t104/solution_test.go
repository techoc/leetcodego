package t104

import (
	"github.com/techoc/leetcodego/leetcode/utils"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		root *utils.TreeNode
		want int
	}{
		{
			name: "Case 1",
			root: utils.ListToTree([]any{3, 9, 20, nil, nil, 15, 7}),
			want: 3,
		},
		{
			name: "Case 2",
			root: utils.ListToTree([]any{1, nil, 2}),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

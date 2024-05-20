package t100

import (
	"github.com/techoc/leetcodego/leetcode/utils"
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    *utils.TreeNode
		q    *utils.TreeNode
		want bool
	}{
		{
			name: "Case 1",
			p:    utils.ListToTree([]any{1, 2, 3}),
			q:    utils.ListToTree([]any{1, 2, 3}),
			want: true,
		},
		{
			name: "Case 2",
			p:    utils.ListToTree([]any{1, 2}),
			q:    utils.ListToTree([]any{1, nil, 2}),
			want: false,
		},
		{
			name: "Case 3",
			p:    utils.ListToTree([]any{1, 2, 1}),
			q:    utils.ListToTree([]any{1, 1, 2}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.p, tt.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

package t226

import "testing"

func TestInvertTree(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	invertTree(root)
	t.Log(root)
}

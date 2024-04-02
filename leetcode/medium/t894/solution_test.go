package t894

import (
	"reflect"
	"testing"
)

func TestAllPossibleFBT(t *testing.T) {
	// 测试用例1: n = 1
	n1 := 1
	expected1 := []*TreeNode{
		&TreeNode{Val: 0},
	}
	result1 := allPossibleFBT(n1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
	}

	// 测试用例2: n = 3
	n2 := 3
	expected2 := []*TreeNode{
		&TreeNode{Val: 0, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}},
	}
	result2 := allPossibleFBT(n2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
	}
}

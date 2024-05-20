package t100

import . "github.com/techoc/leetcodego/leetcode/utils"

// 100. 相同的树
// https://leetcode.cn/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // 两个都为空
		return true
	}
	if p == nil || q == nil { // 一个为空
		return false
	}
	if p.Val != q.Val { // 值不相等
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

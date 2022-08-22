package medium

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	var dfs1 func(node *TreeNode) int
	//递归获取树的高度
	dfs1 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left, right := dfs1(node.Left), dfs1(node.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
	height := dfs1(root)

	//根据高度初始化数组
	n := (1 << height) - 1
	ans := make([][]string, height)
	for i := 0; i < height; i++ {
		ans[i] = make([]string, n)
	}
	var dfs2 func(node *TreeNode, r int, c int)
	dfs2 = func(node *TreeNode, r int, c int) {
		if node == nil {
			return
		}
		ans[r][c] = strconv.Itoa(node.Val)
		if d := height - 2 - r; d >= 0 {
			diff := 1 << d
			dfs2(node.Left, r+1, c-diff)
			dfs2(node.Right, r+1, c+diff)
		}
	}
	dfs2(root, 0, (n-1)>>1)
	return ans
}

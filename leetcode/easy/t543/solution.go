package t543

import (
	. "github.com/techoc/leetcodego/leetcode/utils"
)

// 543. 二叉树的直径
// https://leetcode.cn/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
	var res = 0                   // 初始化结果变量为0
	var depth func(*TreeNode) int // 声明一个内部函数，用于计算节点的深度

	// 定义depth函数，用于递归计算二叉树的深度并更新结果
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0 // 如果节点为空，深度为0
		}
		l := depth(node.Left)  // 计算左子树的深度
		r := depth(node.Right) // 计算右子树的深度
		res = max(res, l+r)    // 更新结果为左右子树深度之和的最大值
		return max(l, r) + 1   // 返回节点的深度，即左右子树最大深度加一
	}
	depth(root) // 从根节点开始计算深度
	return res  // 返回结果
}

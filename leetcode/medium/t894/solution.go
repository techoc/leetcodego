package t894

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var f = [11][]*TreeNode{1: {{}}}

func init() {
	for i := 2; i < len(f); i++ { // 计算 f[i]
		for j := 1; j < i; j++ { // 枚举左子树叶子数
			for _, left := range f[j] { // 枚举左子树
				for _, right := range f[i-j] { // 枚举右子树
					f[i] = append(f[i], &TreeNode{0, left, right})
				}
			}
		}
	}
}

// 894. 所有可能的真二叉树
// https://leetcode.cn/problems/all-possible-full-binary-trees
func allPossibleFBT(n int) []*TreeNode {
	if n%2 > 0 {
		return f[(n+1)/2]
	}
	return nil
}

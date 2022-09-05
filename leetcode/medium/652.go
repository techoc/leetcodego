package medium

import "fmt"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := map[*TreeNode]struct{}{}
	seen := map[string]*TreeNode{}
	var dfs func(*TreeNode) string
	//将树序列化成字符串
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if n, ok := seen[serial]; ok {
			repeat[n] = struct{}{}
		} else {
			seen[serial] = node
		}
		return serial
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}

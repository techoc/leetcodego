package medium

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 1
	q := []pair{{root, 1}}
	for q != nil {
		ans = max(ans, q[len(q)-1].index-q[0].index+1)
		tmp := q
		q = nil
		for _, p := range tmp {
			if p.node.Left != nil {
				q = append(q, pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				q = append(q, pair{p.node.Right, p.index*2 + 1})
			}
		}
	}
	return ans
}

package medium

func constructMaximumBinaryTree(nums []int) *TreeNode {
	//递归终止条件
	if len(nums) == 0 {
		return nil
	}
	best := 0
	for i, num := range nums {
		//找到区域最大值
		if num > nums[best] {
			best = i
		}
	}
	//递归
	return &TreeNode{nums[best], constructMaximumBinaryTree(nums[:best]), constructMaximumBinaryTree(nums[best+1:])}
}

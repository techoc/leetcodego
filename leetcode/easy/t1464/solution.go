package t1464

// 1464. 数组中两元素的最大乘积
// https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/
func maxProduct(nums []int) int {
	// 找到两个最大值
	max1, max2 := nums[0], nums[1]
	if max1 < max2 { // 如果最大值小于次最大值 交换
		max1, max2 = max2, max1
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > max1 { // 如果当前值大于最大值 更新次最大值 更新最大值
			max2 = max1
			max1 = nums[i]
		} else if nums[i] > max2 { // 如果当前值大于次最大值 更新次最大值
			max2 = nums[i]
		}
	}
	return (max1 - 1) * (max2 - 1)
}

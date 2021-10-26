package leetcode

//最大子序和
//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
// 输出：6
// 解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

func MaxSubArray(nums []int) int {
	//贪心
	max := -2147483648
	sum := 0
	for _, value := range nums {
		sum = value + sum
		if sum >= max {
			max = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return max
}

func MaxSubArray1(nums []int) int {
	//动态规划
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

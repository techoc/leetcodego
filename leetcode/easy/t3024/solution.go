package t3024

// 3024.三角形类型
// https://leetcode.cn/problems/type-of-triangle/
func triangleType(nums []int) string {
	// 任意两边之和大于第三边
	if nums[0]+nums[1] > nums[2] && nums[1]+nums[2] > nums[0] && nums[2]+nums[0] > nums[1] {
		// 等腰三角形
		if nums[0] == nums[1] || nums[1] == nums[2] || nums[2] == nums[0] {
			// 等边三角形
			if nums[0] == nums[1] && nums[1] == nums[2] {
				return "equilateral"
			}
			return "isosceles"
		}
		return "scalene"
	}
	return "none"
}

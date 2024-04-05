package t1929

// 1929.数组串联
// https://leetcode.cn/problems/concatenation-of-array
func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

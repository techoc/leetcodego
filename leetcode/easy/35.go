package easy

//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//请必须使用时间复杂度为 O(log n) 的算法。
//示例 1:
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//示例 2:
//输入: nums = [1,3,5,6], target = 2
//输出: 1

//普通遍历 O(n)
func searchInsert(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
		if value > target {
			return index
		}
	}
	//如果没有找到，则返回最后一个元素的下标
	return len(nums)
}

//二分查找 O(log n)
func searchInsertQuick(nums []int, target int) int {
	//初始化左右指针
	left, right := 0, len(nums)-1
	ans := len(nums)
	for left <= right {
		//防止数字大小溢出，等同于(right+left)/2
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

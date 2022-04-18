package medium

//给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//请你找出并返回只出现一次的那个数。
//你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2

//暴力法

func SingleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for k, v := range nums {
		//当重复字符是nums[0]时
		if k == 0 {
			if v != nums[k+1] {
				return v
			} else {
				continue
			}
		} else if k == len(nums)-1 { //如果是最后一位
			return v
		} else if v == nums[k-1] { //如果相等
			continue
		} else if v == nums[k+1] { //如果相等
			continue
		} else {
			return v
		}
	}
	return 0
}

//偶数下标的二分法查找
func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}

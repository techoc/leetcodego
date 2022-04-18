package easy

import "sort"

//	给定一个整数数组，判断是否存在重复元素。
//	如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
///	示例 1:
//	输入: [1,2,3,1]
//	输出: true

func ContainsDuplicate(nums []int) bool {
	//暴力双循环
	//时间复杂度：O(N²)，其中 N 为数组的长度
	//空间复杂度：O(N)，其中 N 为数组的长度
	for index, value := range nums {
		for index1, value1 := range nums {
			if value1 == value && index != index1 {
				return true
			}
		}
	}
	return false
}

func ContainsDuplicateShort(nums []int) bool {
	//在对数字从小到大排序之后，数组的重复元素一定出现在相邻位置中。
	//因此，我们可以扫描已排序的数组，每次判断相邻的两个元素是否相等，如果相等则说明存在重复的元素。
	//时间复杂度：O(N\log N)，其中 N 为数组的长度。需要对数组进行排序。
	//空间复杂度：O(\log N)，其中 N 为数组的长度。注意我们在这里应当考虑递归调用栈的深度。

	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func ContainsDuplicateHashMap(nums []int) bool {
	//对于数组中每个元素，我们将它插入到哈希表中。
	//如果插入一个元素时发现该元素已经存在于哈希表中，则说明存在重复的元素。
	//时间复杂度：O(N)，其中 N 为数组的长度。
	//空间复杂度：O(N)，其中 N 为数组的长度。
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

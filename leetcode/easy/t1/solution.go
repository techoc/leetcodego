package t1

//	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
//	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//	你可以按任意顺序返回答案。

// twoSum 暴力枚举实现
func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// twoSumHashMap 两数之和HashMap实现
func twoSumHashMap(nums []int, target int) []int {
	//	注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。
	//	因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。
	//	使用哈希表，可以将寻找 target - x 的时间复杂度降低到从 O(N) 降低到 O(1)。
	//	这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，
	//	即可保证不会让 x 和自己匹配。
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

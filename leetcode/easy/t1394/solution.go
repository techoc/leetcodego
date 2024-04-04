package t1394

// 1394. 找出数组中的幸运数
// https://leetcode.cn/problems/find-lucky-integer-in-an-array/
func findLucky(arr []int) int {
	// 统计每个数字出现的次数
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}

	res := -1
	for key, value := range count {
		if key == value {
			res = max(res, value)
		}
	}

	return res
}

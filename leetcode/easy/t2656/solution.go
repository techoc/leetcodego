package t2656

func maximizeSum(nums []int, k int) (res int) {
	// 找到nums中的最大值
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	// 模拟
	for i := 0; i < k; i++ {
		res += m
		res += i
	}

	return res
}

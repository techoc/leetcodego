package t2460

func applyOperations(nums []int) []int {
	n := len(nums)
	// 模拟
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] <<= 1
			nums[i+1] = 0
		}
	}
	ans := make([]int, n)
	i := 0
	// 将原数组复制到新数组
	for _, x := range nums {
		if x > 0 {
			ans[i] = x
			i++
		}
	}
	return ans
}

func applyOperationsOffice(nums []int) []int {

	n := len(nums)

	/* 纯模拟 */
	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] <<= 1
			nums[i+1] = 0
		}
	}

	/* 双指针 */
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for ; j < n; j++ {
		nums[j] = 0
	}

	return nums

}

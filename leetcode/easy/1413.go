package easy

func minStartValue(nums []int) int {
	//取最小值
	minNum := 101
	for i, j := 0, 1; i < len(nums); i++ {
		j = j + nums[i]
		minNum = min(minNum, j)
	}
	if minNum < 1 {
		return 2 - minNum
	}
	return 1
}

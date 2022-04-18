package medium

//给你一个整数数组 nums 。nums 中，子数组的 范围 是子数组中最大元素和最小元素的差值。
//返回 nums 中 所有 子数组范围的 和 。
//子数组是数组中一个连续 非空 的元素序列。
//示例 1：
//输入：nums = [1,2,3]
//输出：4
//解释：nums 的 6 个子数组如下所示：
//[1]，范围 = 最大 - 最小 = 1 - 1 = 0
//[2]，范围 = 2 - 2 = 0
//[3]，范围 = 3 - 3 = 0
//[1,2]，范围 = 2 - 1 = 1
//[2,3]，范围 = 3 - 2 = 1
//[1,2,3]，范围 = 3 - 1 = 2
//所有范围的和是 0 + 0 + 0 + 1 + 1 + 2 = 4

func subArrayRanges(nums []int) (ans int64) {
	//遍历子字符串。
	//首先枚举子数组的左边界i 然后枚举子数组的右边界 j并且 i<=j
	//在枚举 j 的过程中我们可以迭代地计算子数组 [i,j] 的最小值与最大值 然后将范围值 最大值-最小值 加到总范围和。
	for i, num := range nums {
		minVal, maxVal := num, num
		for _, v := range nums[i+1:] {
			if v < minVal {
				minVal = v
			} else if v > maxVal {
				maxVal = v
			}
			ans += int64(maxVal - minVal)
		}
	}
	return
}

func solve(nums []int) int64 {
	n := len(nums)
	left := make([]int, n) // left[i] 为左侧严格大于 num[i] 的最近元素位置（不存在时为 -1）
	type pair struct{ v, i int }
	s := []pair{{2e9, -1}} // 哨兵
	for i, v := range nums {
		for s[len(s)-1].v <= v {
			s = s[:len(s)-1]
		}
		left[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	right := make([]int, n) // right[i] 为右侧大于等于 num[i] 的最近元素位置（不存在时为 n）
	s = []pair{{2e9, n}}
	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for s[len(s)-1].v < v {
			s = s[:len(s)-1]
		}
		right[i] = s[len(s)-1].i
		s = append(s, pair{v, i})
	}

	ans := 0
	for i, v := range nums {
		ans += (i - left[i]) * (right[i] - i) * v
	}
	return int64(ans)
}

func SubArrayRanges(nums []int) int64 {
	//单调栈
	ans := solve(nums)
	for i, v := range nums { // 小技巧：所有元素取反后算的就是最小值的贡献
		nums[i] = -v
	}
	return ans + solve(nums)
}

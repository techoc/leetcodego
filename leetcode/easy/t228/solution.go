package t228

import (
	"fmt"
)

func summaryRanges(nums []int) (ans []string) {
	f := func(left, right int) string {
		//如果左右指针相等 则直接添加到结果中
		if left == right {
			return fmt.Sprintf("%v", nums[left])
		}
		//否则添加中间的值 "->"
		return fmt.Sprintf("%v->%v", nums[left], nums[right])
	}

	//双指针遍历整个数组
	for left, right, n := 0, 0, len(nums); left < n; left = right + 1 {
		//初始化左右指针
		right = left

		//如果右指针+1小于数组长度 并且右指针右边的数和右指针是连续的 则右指针右移
		for right+1 < n && nums[right+1] == nums[right]+1 {
			right++
		}

		//调用检查方法 添加结果到结果集中
		ans = append(ans, f(left, right))
	}
	return
}

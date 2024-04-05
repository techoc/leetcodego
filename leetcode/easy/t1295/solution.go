package t1295

import "strconv"

// 1295. 统计位数为偶数的数字个数
// https://leetcode.cn/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			res++
		}
	}
	return res
}

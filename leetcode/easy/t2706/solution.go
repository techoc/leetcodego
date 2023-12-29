package t2706

import "sort"

// 2706. 购买两块巧克力
// https://leetcode.cn/problems/buy-two-chocolates
func buyChoco(prices []int, money int) int {
	// 找到最小的两个数
	sort.Ints(prices)
	if money-prices[0]-prices[1] < 0 {
		return money
	}
	return money - prices[0] - prices[1]
}

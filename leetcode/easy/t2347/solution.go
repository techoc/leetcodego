package t2347

import (
	"bytes"
	"sort"
)

// https://leetcode.cn/problems/best-poker-hand/
func bestHand(ranks []int, suits []byte) string {
	// 判断同花
	if suits[0] == suits[1] && suits[0] == suits[2] && suits[0] == suits[3] && suits[0] == suits[4] {
		return "Flush"
	}
	sort.Ints(ranks)
	// 三条
	if ranks[0] == ranks[2] || ranks[1] == ranks[3] || ranks[2] == ranks[4] {
		return "Three of a Kind"
	}
	//对子
	if ranks[0] == ranks[1] || ranks[1] == ranks[2] || ranks[2] == ranks[3] || ranks[3] == ranks[4] {
		return "Pair"
	}
	//高牌
	return "High Card"
}

func bestHandOffice(ranks []int, suits []byte) string {
	//同花
	if bytes.Count(suits, suits[:1]) == 5 {
		return "Flush"
	}
	cnt, pair := map[int]int{}, false
	// 三条
	for _, r := range ranks {
		cnt[r]++
		if cnt[r] == 3 {
			return "Three of a Kind"
		}
		if cnt[r] == 2 {
			pair = true
		}
	}
	// 对子
	if pair {
		return "Pair"
	}
	// 高牌
	return "High Card"
}

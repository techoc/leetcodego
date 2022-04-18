package medium

import "math"

//给你一个整数数组 bloomDay，以及两个整数 m 和 k 。
//现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。
//花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。
//请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1 。
//示例 1：
//输入：bloomDay = [1,10,3,10,2], m = 3, k = 1
//输出：3
//解释：让我们一起观察这三天的花开过程，x 表示花开，而 _ 表示花还未开。
//现在需要制作 3 束花，每束只需要 1 朵。
//1 天后：[x, _, _, _, _]   只能制作 1 束花
//2 天后：[x, _, _, _, x]    只能制作 2 束花
//3 天后：[x, _, x, _, x]    可以制作 3 束花，答案为 3

func minDays(bloomDay []int, m int, k int) int {
	//特殊返回
	if m*k > len(bloomDay) {
		return -1
	}
	//初始化最大天数和最小天数
	maxDay, minDay := 0, math.MaxInt32
	for _, v := range bloomDay {
		if v > maxDay {
			maxDay = v
		}
		if v < minDay {
			minDay = v
		}
	}

	//二分查找
	for minDay < maxDay {
		mid := (minDay + maxDay) / 2
		//如果可以制作m束花，则最小天数为mid
		if check(bloomDay, m, k, mid) {
			maxDay = mid
		} else {
			//否则最小天数为mid+1
			minDay = mid + 1
		}
	}
	return minDay
}

//检查是否可以制作m束花
func check(bloomDay []int, m int, k int, days int) bool {
	//记录花的数量
	flowers := 0
	//记录花束的数量
	bouquets := 0
	for i := 0; i < len(bloomDay); i++ {
		//
		if bloomDay[i] <= days {
			flowers++
			//如果花开的数量可以做成花束
			if flowers == k {
				bouquets++
				flowers = 0
			}
		} else {
			//如果花开的数量不可以做成花束
			flowers = 0
		}
	}
	return bouquets >= m
}

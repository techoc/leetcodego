package t2363

import "sort"

func mergeSimilarItems(item1, item2 [][]int) [][]int {
	mp := map[int]int{}
	// 遍历item1 将物品的价值 和重量对应
	for _, a := range item1 {
		mp[a[0]] += a[1]
	}
	// 遍历item2 将物品的价值 和重量对应
	for _, a := range item2 {
		mp[a[0]] += a[1]
	}
	var ans [][]int
	for a, b := range mp {
		ans = append(ans, []int{a, b})
	}
	// 按价值升序排序
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

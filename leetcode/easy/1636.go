package easy

import "sort"

func frequencySort(nums []int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return cnt[a] < cnt[b] || cnt[a] == cnt[b] && a > b
	})
	return nums
}

package t2341

// https://leetcode.cn/problems/maximum-number-of-pairs-in-array/
func numberOfPairs(nums []int) []int {
	//
	cnt := [101]int{}
	for _, x := range nums {
		cnt[x]++
	}
	s := 0
	for _, v := range cnt {
		s += v / 2
	}
	return []int{s, len(nums) - s*2}
}

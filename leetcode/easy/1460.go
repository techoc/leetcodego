package easy

import "sort"

func canBeEqual(target, arr []int) bool {
	//初始化哈希表
	cnt := map[int]int{}
	for i, x := range target {
		cnt[x]++
		cnt[arr[i]]--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}

func canBeEqualSort(target, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i, x := range target {
		if x != arr[i] {
			return false
		}
	}
	return true
}

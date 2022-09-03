package medium

import "sort"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func findLongestChain(pairs [][]int) (ans int) {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	cur := INT_MIN
	for _, pair := range pairs {
		if cur < pair[0] {
			cur = pair[1]
			ans++
		}
	}
	return
}

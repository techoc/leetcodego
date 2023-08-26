package t228

import "testing"

func TestSummaryRanges(t *testing.T) {
	t.Log(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	t.Log(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	t.Log(summaryRanges([]int{0, 1}))
	t.Log(summaryRanges([]int{0}))
	t.Log(summaryRanges([]int{1, 3, 6, 9}))
}

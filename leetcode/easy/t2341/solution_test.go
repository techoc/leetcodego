package t2341

import "testing"

func TestNumberOfPairs(t *testing.T) {
	nums := []int{1, 3, 2, 1, 3, 2, 2}
	pairs := numberOfPairs(nums)

	t.Log(pairs)
}

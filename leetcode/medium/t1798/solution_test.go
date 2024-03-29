package t1798

import "testing"

func TestGetMaximumConsecutive(t *testing.T) {
	coins := []int{1, 3}
	consecutive := getMaximumConsecutive(coins)
	t.Log(consecutive)
}

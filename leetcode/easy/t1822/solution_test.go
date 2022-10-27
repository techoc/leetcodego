package t1822

import "testing"

func TestArraySign(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	res := arraySign(nums)
	t.Log(res)
}

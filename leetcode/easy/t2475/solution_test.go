package t2475

import "testing"

func TestUnequalTriplets(t *testing.T) {
	nums := []int{4, 4, 2, 4, 3}
	triplets := unequalTriplets(nums)
	t.Log(triplets)
}

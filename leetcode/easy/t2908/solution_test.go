package t2908

import "testing"

func TestMinimumSum(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{[]int{8, 6, 1, 5, 3}, 9},
		{[]int{5, 4, 8, 7, 10, 2}, 13},
		{[]int{6, 5, 4, 3, 4, 5}, -1},
	}

	for _, test := range tests {
		if ans := minimumSum(test.nums); ans != test.expect {
			t.Errorf("minimumSum(%v) = %d, expect %d", test.nums, ans, test.expect)
		}
	}
}

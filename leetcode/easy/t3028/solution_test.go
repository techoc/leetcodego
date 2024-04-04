package t3028

import "testing"

func TestReturnToBoundaryCount(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "test1",
			nums:     []int{2, 3, -5},
			expected: 1,
		},
		{
			name:     "test2",
			nums:     []int{3, 2, -3, -4},
			expected: 0,
		},
		{
			name:     "test3",
			nums:     []int{2, -2, 2, -2},
			expected: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := returnToBoundaryCount(test.nums); got != test.expected {
				t.Errorf("returnToBoundaryCount(%v) = %v, want %v", test.nums, got, test.expected)
			}
		})
	}
}

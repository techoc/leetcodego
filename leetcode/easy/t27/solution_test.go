package t27

import "testing"

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		val    int
		expect int
	}{
		{
			name:   "case 1",
			nums:   []int{3, 2, 2, 3},
			val:    3,
			expect: 2,
		},
		{
			name:   "case 2",
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			expect: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.nums, tt.val); got != tt.expect {
				t.Errorf("removeElement() = %v, want %v", got, tt.expect)
			}
		})
	}
}

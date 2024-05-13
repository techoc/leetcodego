package t283

import "testing"

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Case 1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "Case 2",
			nums: []int{0},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			for i := 0; i < len(tt.nums); i++ {
				if tt.nums[i] != tt.want[i] {
					t.Errorf("moveZeroes() = %v, want %v", tt.nums, tt.want)
				}
			}
		})
	}
}

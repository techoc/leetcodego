package t1351

import "testing"

func TestCountNegatives(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want int
	}{
		{
			name: "Case 1",
			nums: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			want: 8,
		},
		{
			name: "Case 2",
			nums: [][]int{
				{3, 2},
				{1, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNegatives(tt.nums); got != tt.want {
				t.Errorf("countNegatives() = %v, want %v", got, tt.want)
			}
		})
	}
}

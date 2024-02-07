package t2765

import "testing"

func TestAlternatingSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 3, 4, 3, 4},
			want: 4,
		},
		{
			nums: []int{4, 5, 6},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := alternatingSubarray(tt.nums)
		if got != tt.want {
			t.Errorf("alternatingSubarray() = %v, want %v", got, tt.want)
		}
	}
}

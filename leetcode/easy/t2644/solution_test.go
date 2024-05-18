package t2644

import "testing"

func TestMaxDivScore(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		divisors []int
		want     int
	}{
		{
			name:     "Case 1",
			nums:     []int{4, 7, 9, 3, 9},
			divisors: []int{5, 2, 3},
			want:     3,
		},
		{
			name:     "Case 2",
			nums:     []int{20, 14, 21, 10},
			divisors: []int{5, 7, 5},
			want:     5,
		},
		{
			name:     "Case 3",
			nums:     []int{12},
			divisors: []int{10, 16},
			want:     10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDivScore(tt.nums, tt.divisors); got != tt.want {
				t.Errorf("maxDivScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

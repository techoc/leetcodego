package t1464

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Case 1", []int{3, 4, 5, 2}, 12},
		{"Case 2", []int{1, 5, 4, 5}, 16},
		{"Case 3", []int{3, 7}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

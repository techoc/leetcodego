package t3024

import "testing"

func TestTriangleType(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{
			name: "Case 1",
			nums: []int{3, 3, 3},
			want: "equilateral",
		},
		{
			name: "Case 2",
			nums: []int{3, 4, 5},
			want: "scalene",
		},
		{
			name: "Case 3",
			nums: []int{4, 4, 6},
			want: "isosceles",
		},
		{
			name: "Case 4",
			nums: []int{10, 1, 1},
			want: "none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}

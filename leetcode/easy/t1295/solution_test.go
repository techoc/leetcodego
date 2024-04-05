package t1295

import "testing"

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case 1",
			nums: []int{12, 345, 2, 6, 7896},
			want: 2,
		},
		{
			name: "Case 2",
			nums: []int{555, 901, 482, 1771},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

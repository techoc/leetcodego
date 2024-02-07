package t215

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		args []int
		k    int
		want int
	}{
		{
			name: "test-1",
			args: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "test-2",
			args: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args, tt.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

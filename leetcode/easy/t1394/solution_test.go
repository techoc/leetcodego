package t1394

import "testing"

func TestFindLucky(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"test1", []int{2, 2, 3, 4}, 2},
		{"test2", []int{1, 2, 2, 3, 3, 3}, 3},
		{"test3", []int{2, 2, 2, 3, 3}, -1},
		{"test4", []int{2, 2, 2, 3, 3}, -1},
		{"test5", []int{5, 8, 10, 9, 8, 10, 9, 6, 6, 7, 10, 8, 10, 10, 9, 4, 6, 2, 10, 3, 5, 10, 2, 6, 4, 8, 7, 3, 9, 9, 5, 7, 8, 7, 11, 9, 3, 1, 5, 11, 9, 5, 8, 10, 8, 4, 9, 7, 6, 7, 10, 9, 7, 8, 6, 10, 4}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.args); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}

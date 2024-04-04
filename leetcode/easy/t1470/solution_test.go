package t1470

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		n    int
		want []int
	}{
		{
			"Case 1",
			[]int{2, 5, 1, 3, 4, 7},
			3,
			[]int{2, 3, 5, 4, 1, 7},
		},
		{
			"Case 2",
			[]int{1, 2, 3, 4, 4, 3, 2, 1},
			4,
			[]int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			"Case 3",
			[]int{1, 1, 2, 2},
			2,
			[]int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.nums, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

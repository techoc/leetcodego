package t1103

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		name       string
		candies    int
		num_people int
		want       []int
	}{
		{
			name:       "Case 1",
			candies:    7,
			num_people: 4,
			want:       []int{1, 2, 3, 1},
		},
		{
			name:       "Case 2",
			candies:    10,
			num_people: 3,
			want:       []int{5, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.candies, tt.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}

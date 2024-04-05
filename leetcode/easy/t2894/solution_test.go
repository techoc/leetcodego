package t2894

import "testing"

func TestDifferenceOfSums(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int
	}{
		{
			name: "Case 1",
			n:    10,
			m:    3,
			want: 19,
		},
		{
			name: "Case 2",
			n:    5,
			m:    6,
			want: 15,
		}, {
			name: "Case 3",
			n:    5,
			m:    1,
			want: -15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSums(tt.n, tt.m); got != tt.want {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

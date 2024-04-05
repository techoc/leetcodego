package t2798

import "testing"

func TestNumberOfEmployeesWhoMetTarget(t *testing.T) {
	tests := []struct {
		name   string
		hours  []int
		target int
		want   int
	}{
		{
			name:   "Case 1",
			hours:  []int{0, 1, 2, 3, 4},
			target: 2,
			want:   3,
		},
		{
			name:   "Case 2",
			hours:  []int{5, 1, 4, 2, 2},
			target: 6,
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfEmployeesWhoMetTarget(tt.hours, tt.target); got != tt.want {
				t.Errorf("numberOfEmployeesWhoMetTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

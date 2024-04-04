package t1450

import "testing"

func TestBusyStudent(t *testing.T) {
	tests := []struct {
		name      string
		startTime []int
		endTime   []int
		queryTime int
		want      int
	}{
		{"Case 1", []int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{"Case 2", []int{4}, []int{4}, 4, 1},
		{"Case 3", []int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7, 0},
		{"Case 4", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := busyStudent(tt.startTime, tt.endTime, tt.queryTime); got != tt.want {
				t.Errorf("busyStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}

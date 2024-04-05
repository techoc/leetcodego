package t1137

import "testing"

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Case 1", 4, 4},
		{"Case 2", 25, 1389537},
		{"Case 2", 35, 615693474},
		{"Case 3", 0, 0},
		{"Case 4", 1, 1},
		{"Case 5", 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

package t924

import "testing"

func TestMinMalwareSpread(t *testing.T) {
	tests := []struct {
		name    string
		graph   [][]int
		initial []int
		want    int
	}{
		{"Case 1", [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, []int{0, 1}, 0},
		{"Case 2", [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}, []int{0, 2}, 0},
		{"Case 3", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, []int{1, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMalwareSpread(tt.graph, tt.initial); got != tt.want {
				t.Errorf("minMalwareSpread() = %v, want %v", got, tt.want)
			}
		})
	}
}

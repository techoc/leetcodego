package t100256

import "testing"

func TestFindLatestTime(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"Case 1", "1?:?4", "11:54"},
		{"Case 2", "0?:5?", "09:59"},
		{"Case 3", "?1:5?", "11:59"},
		{"Case 4", "00:28", "00:28"},
		{"Case 5", "00:4?", "00:49"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLatestTime(tt.s); got != tt.want {
				t.Errorf("findLatestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

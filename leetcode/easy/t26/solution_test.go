package t26

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "case 1",
			input: []int{1, 1, 2},
			want:  2,
		},
		{
			name:  "case 2",
			input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.input); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

package t2769

import "testing"

func TestTheMaximumAchievableX(t *testing.T) {
	tests := []struct {
		name string
		num  int
		t    int
		want int
	}{
		{
			name: "Case 1",
			num:  4,
			t:    1,
			want: 6,
		},
		{
			name: "Case 2",
			num:  3,
			t:    2,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := theMaximumAchievableX(tt.num, tt.t); got != tt.want {
				t.Errorf("theMaximumAchievableX() = %v, want %v", got, tt.want)
			}
		})
	}
}

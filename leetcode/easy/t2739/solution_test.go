package t2739

import "testing"

func TestDistanceTraveled(t *testing.T) {
	tests := []struct {
		name           string
		mainTank       int
		additionalTank int
		want           int
	}{
		{name: "Case 1", mainTank: 1, additionalTank: 2, want: 10},
		{name: "Case 2", mainTank: 5, additionalTank: 10, want: 60},
		{name: "Case 3", mainTank: 9, additionalTank: 2, want: 110},
		{name: "Case 4", mainTank: 10, additionalTank: 2, want: 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceTraveled(tt.mainTank, tt.additionalTank); got != tt.want {
				t.Errorf("distanceTraveled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceTraveledTwo(t *testing.T) {
	tests := []struct {
		name           string
		mainTank       int
		additionalTank int
		want           int
	}{
		{name: "Case 1", mainTank: 1, additionalTank: 2, want: 10},
		{name: "Case 2", mainTank: 5, additionalTank: 10, want: 60},
		{name: "Case 3", mainTank: 9, additionalTank: 2, want: 110},
		{name: "Case 4", mainTank: 10, additionalTank: 2, want: 120},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceTraveledTwo(tt.mainTank, tt.additionalTank); got != tt.want {
				t.Errorf("distanceTraveled() = %v, want %v", got, tt.want)
			}
		})
	}
}

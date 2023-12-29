package t2706

import "testing"

func TestBuyChoco(t *testing.T) {
	tests := []struct {
		prices []int
		money  int
		want   int
	}{
		{[]int{1, 2, 2}, 3, 0},
		{[]int{3, 5, 7}, 15, 7},
		{[]int{1, 2, 3}, 5, 2},
		{[]int{5, 10, 15}, 20, 5},
	}

	for _, test := range tests {
		got := buyChoco(test.prices, test.money)
		if got != test.want {
			t.Errorf("buyChoco(%v, %d) = %d, want %d", test.prices, test.money, got, test.want)
		}
	}
}

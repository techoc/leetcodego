package t2303

import "testing"

func TestCalculateTax(t *testing.T) {
	var brackets [][]int

	bracket1 := []int{1, 0}
	bracket2 := []int{4, 25}
	bracket3 := []int{5, 510}

	brackets = append(brackets, bracket1)
	brackets = append(brackets, bracket2)
	brackets = append(brackets, bracket3)
	tax := calculateTax(brackets, 2)
	tax1 := calculateTax1(brackets, 2)

	t.Log(tax)
	t.Log(tax1)
}

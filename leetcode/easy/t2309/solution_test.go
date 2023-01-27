package t2309

import "testing"

func TestGreatestLetter(t *testing.T) {
	s := "lEeTcOdE"
	letter := greatestLetter(s)
	t.Log(letter)
}

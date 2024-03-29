package t1790

import (
	"log"
	"testing"
)

func TestAreAlmostEqual(t *testing.T) {
	s1 := "aa"
	s2 := "ac"

	result := areAlmostEqual(s1, s2)

	log.Println(result)
}

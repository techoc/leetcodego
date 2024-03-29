package t2451

import (
	"log"
	"testing"
)

func TestOddString(t *testing.T) {
	words := []string{"adc", "wzy", "abc"}
	s := oddString(words)
	log.Print(s)
}

package t415

import (
	"log"
	"testing"
)

func TestAddStrings(t *testing.T) {
	str1 := "123"
	str2 := "123"
	strings := addStrings(str1, str2)

	log.Print(strings)
}

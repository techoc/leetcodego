package t1668

import "testing"

func TestMaxRepeating(t *testing.T) {
	sequence := "ababc"
	words := "ab"
	ans := maxRepeating(sequence, words)
	t.Log(ans)
}

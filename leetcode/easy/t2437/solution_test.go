package t2437

import "testing"

func TestCountTime(t *testing.T) {
	time := "?5:00"
	i := countTime(time)
	t.Log(i)
}

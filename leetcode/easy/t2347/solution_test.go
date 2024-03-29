package t2347

import "testing"

func TestBestHand(t *testing.T) {
	ranks := []int{9, 2, 13, 1, 2}
	suits := []byte{'b', 'd', 'd', 'b', 'c'}
	hand := bestHand(ranks, suits)
	t.Log(hand)
	t.Log(bestHandOffice(ranks, suits))

}

package t705

import "testing"

func TestMyHashSet(t *testing.T) {
	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	if !myHashSet.Contains(1) {
		t.Error("myHashSet.Contains(1) = false, want true")
	}
	if myHashSet.Contains(3) {
		t.Error("myHashSet.Contains(3) = true, want false")
	}
	myHashSet.Add(2)
	if !myHashSet.Contains(2) {
		t.Error("myHashSet.Contains(2) = false, want true")
	}
	myHashSet.Remove(2)
	if myHashSet.Contains(2) {
		t.Error("myHashSet.Contains(2) = true, want false")
	}

}

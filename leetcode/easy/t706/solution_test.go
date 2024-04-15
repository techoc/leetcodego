package t706

import "testing"

func TestMyHashMap(t *testing.T) {
	myHashMap := Constructor()
	myHashMap.Put(1, 1)
	myHashMap.Put(2, 2)
	if myHashMap.Get(1) != 1 {
		t.Fatal("error")
	}
	if myHashMap.Get(3) != -1 {
		t.Fatal("error")
	}
	myHashMap.Put(2, 1)
	if myHashMap.Get(2) != 1 {
		t.Fatal("error")
	}
	myHashMap.Remove(2)
	if myHashMap.Get(2) != -1 {
		t.Fatal("error")
	}
}

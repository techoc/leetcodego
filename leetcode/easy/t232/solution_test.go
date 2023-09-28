package t232

import "testing"

func TestMyQueue(t *testing.T) {
	obj := Constructor()
	x := 1
	obj.Push(x)
	obj.Push(x)
	obj.Push(x)
	param2 := obj.Pop()
	param3 := obj.Peek()
	param4 := obj.Empty()

	t.Log(param2, param3, param4)
}

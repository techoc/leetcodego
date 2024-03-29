package t1670

import "testing"

func TestFrontMiddleBackQueue(t *testing.T) {
	obj := Constructor()
	obj.PushFront(1)
	obj.PushBack(2)
	obj.PushMiddle(3)
	obj.PushMiddle(4)
	param1 := obj.PopFront()
	param2 := obj.PopMiddle()
	param3 := obj.PopMiddle()
	param4 := obj.PopBack()
	param5 := obj.PopFront()
	t.Log(param1, param2, param3, param4, param5)
}

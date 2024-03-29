package t225

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
}

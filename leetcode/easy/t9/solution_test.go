package t9

import "testing"

func TestIsPalindrome(t *testing.T) {
	x := 121
	palindrome := isPalindrome(x)
	t.Log(palindrome)
}

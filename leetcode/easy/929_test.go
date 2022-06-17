package easy

import "testing"

func Test929(t *testing.T) {
	emails := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	res := numUniqueEmails(emails)
	t.Log(res)
}

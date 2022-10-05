package _811

import (
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	strings := subdomainVisits([]string{"9001 discuss.leetcode.com"})
	for _, s := range strings {
		t.Log(s)
	}
}

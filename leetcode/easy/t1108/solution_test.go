package t1108

import "testing"

func TestDefangIPaddr(t *testing.T) {
	testCases := []struct {
		address  string
		expected string
	}{
		{"192.168.0.1", "192[.]168[.]0[.]1"},
		{"10.0.0.1", "10[.]0[.]0[.]1"},
		{"172.16.254.1", "172[.]16[.]254[.]1"},
	}

	for _, tc := range testCases {
		result := defangIPaddr(tc.address)
		if result != tc.expected {
			t.Errorf("Expected defangIPaddr(%s) to be %s, but got %s", tc.address, tc.expected, result)
		}
	}
}

func TestDefangIPaddr2(t *testing.T) {
	testCases := []struct {
		address  string
		expected string
	}{
		{"192.168.0.1", "192[.]168[.]0[.]1"},
		{"10.0.0.1", "10[.]0[.]0[.]1"},
		{"172.16.254.1", "172[.]16[.]254[.]1"},
	}

	for _, tc := range testCases {
		result := defangIPaddr2(tc.address)
		if result != tc.expected {
			t.Errorf("Expected defangIPaddr2(%s) to be %s, but got %s", tc.address, tc.expected, result)
		}
	}
}

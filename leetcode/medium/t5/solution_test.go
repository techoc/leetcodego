package t5

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Case 1",
			input:    "babad",
			expected: "bab",
		},
		{
			name:     "Case 2",
			input:    "cbbd",
			expected: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.input); got != tt.expected {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.expected)
			}
		})
	}
}

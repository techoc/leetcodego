package t1455

import "testing"

func TestIsPrefixOfWord(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		search   string
		want     int
	}{
		{"Case 1", "i love eating burger", "burg", 4},
		{"Case 2", "this problem is an easy problem", "pro", 2},
		{"Case 3", "i am tired", "you", -1},
		{"Case 4", "i use triple pillow", "pill", 4},
		{"Case 5", "hellohello hellohellohello", "ell", -1},
		{"Case 6", "hello from the other side", "her", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrefixOfWord(tt.sentence, tt.search); got != tt.want {
				t.Errorf("isPrefixOfWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

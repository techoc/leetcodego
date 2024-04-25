package t811

import (
	"github.com/techoc/leetcodego/leetcode/utils"
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	tests := []struct {
		name      string
		cpdomains []string
		want      []string
	}{
		{
			name:      "Case 1",
			cpdomains: []string{"9001 discuss.leetcode.com"},
			want:      []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"},
		},
		{
			name:      "Case 2",
			cpdomains: []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			want:      []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits(tt.cpdomains); !utils.StringListIsEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}

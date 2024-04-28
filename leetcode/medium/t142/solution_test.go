package t142

import (
	. "github.com/techoc/leetcodego/leetcode/utils"
	"reflect"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{
			name: "Case 3",
			head: ListToListNode([]int{1}),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

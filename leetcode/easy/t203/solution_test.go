package t203

import (
	. "github.com/techoc/leetcodego/leetcode/utils"
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		val  int
		want *ListNode
	}{
		{
			name: "Case 1",
			head: ListToListNode([]int{1, 2, 6, 3, 4, 5, 6}),
			val:  6,
			want: ListToListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			name: "Case 2",
			head: ListToListNode([]int{}),
			val:  1,
			want: nil,
		},
		{
			name: "Case 3",
			head: ListToListNode([]int{7, 7, 7, 7}),
			val:  7,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.head, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", ListNodeToList(got), ListNodeToList(tt.want))
			}
		})
	}
}

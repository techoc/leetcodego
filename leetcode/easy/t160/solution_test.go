package t160

import (
	. "github.com/techoc/leetcodego/leetcode/utils"
	"reflect"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name         string
		headA        *ListNode
		headB        *ListNode
		intersectVal int
		skipA        int
		skipB        int
		want         *ListNode
	}{
		{
			name:         "Case 1",
			headA:        ListToListNode([]int{4, 1, 8, 4, 5}),
			headB:        ListToListNode([]int{5, 6, 1, 8, 4, 5}),
			intersectVal: 8,
			skipA:        2,
			skipB:        3,
			want:         ListToListNode([]int{8, 4, 5}),
		},
		{
			name:         "Case 2",
			headA:        ListToListNode([]int{1, 9, 1, 2, 4}),
			headB:        ListToListNode([]int{3, 2, 4}),
			intersectVal: 2,
			skipA:        3,
			skipB:        1,
			want:         ListToListNode([]int{2, 4}),
		},
		{
			name:         "Case 3",
			headA:        ListToListNode([]int{2, 6, 4}),
			headB:        ListToListNode([]int{1, 5}),
			intersectVal: 0,
			skipA:        3,
			skipB:        2,
			want:         nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeListNode(tt.intersectVal, tt.headA, tt.headB, tt.skipA, tt.skipB)
			if got := getIntersectionNode(tt.headA, tt.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeListNode(intersectVal int, listA *ListNode, listB *ListNode, skipA int, skipB int) {
	headA := listA
	for i := 0; i < skipA; i++ {
		headA = headA.Next
	}
	headB := listB
	for i := 0; i < skipB; i++ {
		headB = headB.Next
	}
	if intersectVal == 0 {
		return
	}
	headB.Next = headA
}

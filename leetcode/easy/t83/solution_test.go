package t83

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	listNode := deleteDuplicates(head)
	for listNode != nil {
		t.Logf("%d", listNode.Val)
		listNode = listNode.Next
	}

}

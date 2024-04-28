package listNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToListNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	head := &ListNode{Val: list[0]}
	tail := head
	for i := 1; i < len(list); i++ {
		tail.Next = &ListNode{Val: list[i]}
		tail = tail.Next
	}
	return head
}

func ListNodeToList(list *ListNode) []int {
	if list == nil {
		return []int{}
	}
	var res []int
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return res
}

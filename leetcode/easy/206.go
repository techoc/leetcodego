package easy

func reverseList(head *ListNode) *ListNode {
	//存储上一个节点
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

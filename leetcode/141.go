package leetcode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	//快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		//慢指针走一步
		slow = slow.Next
		//快指针走两布
		fast = fast.Next.Next
		//如果两个指针相遇，则有环
		if slow == fast {
			return true
		}
	}
	return false
}

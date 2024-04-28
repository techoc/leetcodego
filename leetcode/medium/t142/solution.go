package t142

import . "github.com/techoc/leetcodego/leetcode/utils"

// 142.环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	// 快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast { // 快慢指针相遇 则有环 寻找入环点
			break
		}
	}
	if fast == nil || fast.Next == nil { // 没有环
		return nil
	}
	// 重置慢指针
	slow = head
	// 快慢指针再次相遇 则为入环点
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

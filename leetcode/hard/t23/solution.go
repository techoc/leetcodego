package t23

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import . "github.com/techoc/leetcodego/leetcode/utils"

// 23.合并K个排序链表
// https://leetcode.cn/problems/merge-k-sorted-lists
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	for i := 0; i < len(lists); i++ {
		if i == 0 {
			pre = lists[i]
			continue
		}
		cur = lists[i]
		pre = mergeTwoLists(pre, cur) // 合并两个链表
	}
	return pre
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哑节点作为新链表的头节点
	dummy := &ListNode{}
	// tail用于跟踪新链表的当前最后一个节点
	tail := dummy
	// aPtr 和 bPtr 用于遍历链表 l1 和 l2
	aPtr, bPtr := l1, l2
	// 当两个链表都不为空时，循环遍历
	for aPtr != nil && bPtr != nil {
		// 比较当前两个节点的值，将较小的节点添加到新链表中
		if aPtr.Val < bPtr.Val {
			tail.Next = aPtr
			aPtr = aPtr.Next
		} else {
			tail.Next = bPtr
			bPtr = bPtr.Next
		}
		// 移动tail到新链表的下一个节点
		tail = tail.Next
	}

	// 将剩余的节点（如果有的话）连接到新链表的末尾
	tail.Next = aPtr
	if tail.Next == nil {
		tail.Next = bPtr
	}
	// 返回新链表的头节点（跳过虚拟节点）
	return dummy.Next
}

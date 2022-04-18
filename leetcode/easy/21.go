package easy

import "github.com/techoc/leetcodego/leetcode/medium"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]

func MergeTwoLists(list1 *medium.ListNode, list2 *medium.ListNode) *medium.ListNode {
	//特殊情况
	if list1 == nil {
		return list2
	}
	//特殊情况
	if list2 == nil {
		return list1
	}
	//比较大小
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
}

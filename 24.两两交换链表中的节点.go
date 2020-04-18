/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	re := head.Next
	head.Next = swapPairs(head.Next.Next)
	re.Next = head
	return re
}
// @lc code=end


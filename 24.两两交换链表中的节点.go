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
	re := &ListNode{}
	cur := re
	i := 0
	var pre *ListNode
	for head != nil {
		i++
		node := head
		head = head.Next
		if i%2 != 0 {
			pre = node
			continue
		}
		cur.Next = node
		cur = cur.Next
		cur.Next = pre
		cur = cur.Next
	}

	if i%2 != 0 {
		cur.Next = pre
	} else {
		cur.Next = head
	}

	return re.Next
}
// @lc code=end


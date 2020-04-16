package main

import "fmt"

func main() {
	list := []int{}
	l1 := generateList(list)

	dumpList(l1)

	l2 := swapPairs(l1)

	dumpList(l2)
	// dumpList(l3)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func generateList(list []int) *ListNode {
	var head *ListNode
	if len(list) == 0 {
		return head
	}
	head = &ListNode{
		Val: list[0],
	}
	cur := head
	for i := 1; i < len(list); i++ {
		cur.Next = &ListNode{
			Val: list[i],
		}
		cur = cur.Next
	}

	return head
}

func dumpList(list *ListNode) {
	for list != nil {
		fmt.Print("->", list.Val)
		list = list.Next
	}
	fmt.Println()
}

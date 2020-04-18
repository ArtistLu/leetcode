package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
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
	if head == nil || head.Next == nil {
		return head
	}

	re := head.Next
	head.Next = swapPairs(head.Next.Next)
	re.Next = head
	return re
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

package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	var s *ListNode
	for head != nil {
		s, head, head.Next = head, head.Next, s
	}

	return s
}

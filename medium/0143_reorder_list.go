package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	secondHalfStart := slow.Next
	slow.Next = nil
	secondHalf := reverseList(secondHalfStart)
	dummy := &ListNode{}
	current := dummy
	for secondHalf != nil && head != nil {
		currH := head.Next
		currSH := secondHalf.Next
		head.Next = nil
		secondHalf.Next = nil
		current.Next = head
		current.Next.Next = secondHalf
		head = currH
		secondHalf = currSH
		current = current.Next.Next
	}
	if head != nil {
		current.Next = head
	}
	head = dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = next
	}
	return dummy.Next
}

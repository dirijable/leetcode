package medium

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := createStack(l1), createStack(l2)
	carry := 0
	dummy := &ListNode{}
	for len(s1) != 0 || len(s2) != 0 {
		currentSum := 0
		if len(s1) != 0 {
			currentSum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) != 0 {
			currentSum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		currentSum += carry
		carry = 0
		for currentSum >= 10 {
			currentSum -= 10
			carry++
		}
		node := &ListNode{Val: currentSum, Next: dummy.Next}
		dummy.Next = node
		currentSum = 0
	}
    if carry > 0 {
        node := &ListNode{Val: carry, Next: dummy.Next}
		dummy.Next = node
    }
	return dummy.Next

}

func createStack(l *ListNode) []int {
	s := make([]int, 0)
	for tmp := l; tmp != nil; tmp = tmp.Next {
		s = append(s, tmp.Val)
	}
	return s
}
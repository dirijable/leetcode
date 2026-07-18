package easy


type ListNodeR struct {
    Val int
    Next *ListNodeR
}

func reverseList(head *ListNodeR) *ListNodeR {
    dummy := &ListNodeR{}
	for head.Next != nil {
		next := head.Next
		head.Next = dummy.Next
		dummy.Next = head
		head = next
	} 
	return dummy.Next
}

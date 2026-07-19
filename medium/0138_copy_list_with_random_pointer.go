package medium

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	nodes := make(map[*Node]*Node, 0)
	dummy := &Node{}
	curr := dummy
	for head != nil {
		node := getAndPut(nodes, head)
		node.Next = getAndPut(nodes, head.Next)
		node.Random = getAndPut(nodes, head.Random)
		curr.Next = node
        head = head.Next
        curr = curr.Next
	}
	return dummy.Next
}

func getAndPut(nodes map[*Node]*Node, key *Node) *Node {
	var node *Node 
	if n, ok := nodes[key]; ok {
		node = n
	} else if key != nil{
		node = &Node{Val: key.Val}
		nodes[key] = node
	}
	return node
}

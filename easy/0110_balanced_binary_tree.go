package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	balanced := true
    depth(root, &balanced)
	return balanced
}

func depth(root *TreeNode, balanced *bool) int {
	if root == nil {
		return 0
	}
	leftDepth := depth(root.Left, balanced)
	rightDepth := depth(root.Right, balanced)
	mx := max(leftDepth, rightDepth)
 	if mx - min(leftDepth,rightDepth) > 1 {
		*balanced = false
		return -1
	}
	return 1 + mx
}
package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	dfs(root, &maxDiameter)
	return maxDiameter
}

func dfs(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}
	depthL := dfs(root.Left, maxDiameter)
	depthR := dfs(root.Right, maxDiameter)
	
	*maxDiameter = max(depthL+depthR, *maxDiameter)
	
	return 1 + max(depthL, depthR)
}

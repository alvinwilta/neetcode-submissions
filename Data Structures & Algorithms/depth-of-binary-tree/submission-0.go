/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	depthLeft := maxDepth(root.Left)
	depthRight := maxDepth(root.Right)
	if depthLeft > depthRight {
		return depthLeft+1
	}

	return depthRight+1
}
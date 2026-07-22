/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return goodNodeRec(root, root.Val)
}

func goodNodeRec(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}
	gNodes := 0
	
	if root.Val >= max {
		max = root.Val
		gNodes++
	}

	gNodes += goodNodeRec(root.Left,max)
	gNodes += goodNodeRec(root.Right,max)

	return gNodes
}
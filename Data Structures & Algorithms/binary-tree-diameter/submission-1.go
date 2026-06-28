/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var maxLength int
	diameter(&maxLength, root)
	return maxLength
}

func diameter(maxLength *int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftLength, rightLength int
	if root.Left != nil {
	leftLength = diameter(maxLength, root.Left)+1
	}
	if root.Right != nil {
	rightLength = diameter(maxLength, root.Right)+1
	}
	d := leftLength+rightLength
	if d > *maxLength {
		*maxLength = d
	}

	if leftLength > rightLength {
		return leftLength
	}

	return rightLength
}

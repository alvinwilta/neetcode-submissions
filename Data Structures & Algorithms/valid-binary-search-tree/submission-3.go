/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root == nil {
		return true
	}

	return isValidBSTRec(root, 1001, -1001)
}

func isValidBSTRec(root *TreeNode, max, min int) bool {
	leftCheck, rightCheck := true, true 
	if root.Left != nil {
		if root.Left.Val < root.Val && root.Left.Val > min {
			leftCheck = isValidBSTRec(root.Left, root.Val, min)
		} else {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val > root.Val && root.Right.Val < max {
			rightCheck = isValidBSTRec(root.Right, max, root.Val)
		} else {
			return false
		}
	}

	return leftCheck && rightCheck
}

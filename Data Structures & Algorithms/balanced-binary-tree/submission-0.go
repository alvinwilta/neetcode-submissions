/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	
	isBalanced := true
	hLeft := getHeight(&isBalanced, root.Left)
	hRight := getHeight(&isBalanced, root.Right)
	if isBalanced {
		return !(hLeft-hRight > 1 || hRight-hLeft > 1)
	}

	return isBalanced
}

func getHeight(isBalanced *bool, root *TreeNode) int {
	if !*isBalanced {
		return 0
	}

	if root == nil {
		return 0
	}

	hLeft := getHeight(isBalanced, root.Left)
	hRight := getHeight(isBalanced, root.Right)

	if hLeft-hRight > 1 || hRight-hLeft > 1 {
		*isBalanced = false
	}

	if hLeft > hRight {
		return hLeft+1
	}

	return hRight+1
}
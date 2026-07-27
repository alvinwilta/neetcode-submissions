/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) <= 0 || len(inorder) <= 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	// find the root node
	rootIndex := 0
	for i, rootVal := range inorder {
		if rootVal == preorder[0] {
			rootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

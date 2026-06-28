/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    // base case
    if root == nil {
        return nil
    }

    // recursive case
    right := invertTree(root.Left)
    left := invertTree(root.Right)

    root.Right = right
    root.Left = left

    return root
}

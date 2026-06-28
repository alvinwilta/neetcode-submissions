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

    var (
        depthLeft int
        depthRight int
    )

    if root.Left != nil {
        depthLeft = maxDepth(root.Left)
    }

    if root.Right != nil {
        depthRight = maxDepth(root.Right)
    }

    // +1 for the current level
    if depthLeft > depthRight {
        return depthLeft+1
    }

    return depthRight+1
}


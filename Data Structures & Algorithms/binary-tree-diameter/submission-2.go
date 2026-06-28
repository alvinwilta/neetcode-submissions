/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var maxDiam int
    diameter(root, &maxDiam)

    return maxDiam
}

func diameter(root *TreeNode, maxDiam *int) int {
    if root == nil {
        return 0
    }

    var (
        lengthLeft int
        lengthRight int
    )

    if root.Left != nil {
        lengthLeft = diameter(root.Left, maxDiam)+1
    }

    if root.Right != nil {
        lengthRight = diameter(root.Right, maxDiam)+1
    }

    if lengthLeft+lengthRight > *maxDiam {
        *maxDiam = lengthLeft+lengthRight
    }

    if lengthLeft > lengthRight {
        return lengthLeft
    }

    return lengthRight
}

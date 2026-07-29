/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    _, best := chain(root)
	return best
}

func chain(root *TreeNode) (down, best int) {
	if root == nil {
		return 0, -1001
	}

	lDown, lBest := chain(root.Left)
	rDown, rBest := chain(root.Right)

	l,r := max(lDown,0), max(rDown,0) // choice to ignore chain
	return max(l,r)+root.Val, max(max(lBest,rBest), l+r+root.Val) 
}

func max(a,b int) int {
	if a>b{return a}
	return b
}
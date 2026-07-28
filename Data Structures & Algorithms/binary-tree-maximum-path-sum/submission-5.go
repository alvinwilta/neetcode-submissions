/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, max := maxSubPath(root)
	return max
}

func maxSubPath(root *TreeNode) (int, int) {
	if root == nil {
		return -1001,-1001
	}

	leftSum, leftGlobal := maxSubPath(root.Left)
	rightSum, rightGlobal := maxSubPath(root.Right)
	// sum: the sum if we want to go that path
	// e.g. leftSum is the sum of the left path including left root
	// global: the current max of that specific path, detached or not

	// handle local sum
	leftChain := root.Val
	rightChain := root.Val
	if leftSum+leftChain > leftChain {
		leftChain += leftSum
	}
	if rightSum+rightChain > rightChain {
		rightChain += rightSum
	}
	localSum := root.Val
	if localSum < leftChain {
		localSum = leftChain 
	}
	if localSum < rightChain {
		localSum = rightChain
	}

	// handle globals
	localGlobal := leftGlobal
	if rightGlobal > leftGlobal {
		localGlobal = rightGlobal
	}
	if localSum > localGlobal {
		localGlobal = localSum
	}
	if leftSum+rightSum+root.Val > localGlobal {
		localGlobal = leftSum+rightSum+root.Val
	}


	return localSum, localGlobal
}

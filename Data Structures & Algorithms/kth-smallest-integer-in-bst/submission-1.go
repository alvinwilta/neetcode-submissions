/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    if root == nil {
		return 0
	}

	return kthSmallestRec(root, &k) 
}

func kthSmallestRec(root *TreeNode, k *int) int {
	if root.Left != nil {
		leftRec := kthSmallestRec(root.Left, k)
		if *k == 0 {
			return leftRec
		}
	}

	*k -= 1
	if *k == 0 {
		return root.Val
	}

	if root.Right != nil {
		rightRec := kthSmallestRec(root.Right, k)
		if *k == 0 {
			return rightRec
		}
	}

	
	return root.Val
}
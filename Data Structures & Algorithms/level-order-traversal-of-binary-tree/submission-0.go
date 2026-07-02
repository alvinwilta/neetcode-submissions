/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var leftRes, rightRes [][]int

	if root.Left != nil {
		leftRes = levelOrder(root.Left)
	}

	if root.Right != nil {
		rightRes = levelOrder(root.Right)
	}

	return append([][]int{{root.Val}}, mergeArr(leftRes, rightRes)...)
}

func mergeArr(arr1, arr2 [][]int) [][]int {
	res := make([][]int, 0)
	if len(arr1) > len(arr2) {
		for i:=0;i<len(arr2);i++ {
			res = append(res, append(arr1[i],arr2[i]...))
		}
		res = append(res, arr1[len(arr2):]...)
	} else {
		for i:=0;i<len(arr1);i++ {
			res = append(res, append(arr1[i],arr2[i]...))
		}
		res = append(res, arr2[len(arr1):]...)
	}
	return res
}
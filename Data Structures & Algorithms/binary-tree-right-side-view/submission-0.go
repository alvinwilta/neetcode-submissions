/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    //level order traversal
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := make([]*TreeNode,0)
	queue = append(queue, root)
	processQueue(&res, queue)

	return res
}

func processQueue(res *[]int, queue []*TreeNode) {
	*res = append(*res, queue[len(queue)-1].Val) // rightmost
	nextQueue := make([]*TreeNode,0)
	for _, node := range queue {
		if node.Left != nil {
			nextQueue = append(nextQueue, node.Left)
		}
		if node.Right != nil {
			nextQueue = append(nextQueue, node.Right)
		}
	}

	if len(nextQueue) > 0 {
		processQueue(res, nextQueue)
	}
}



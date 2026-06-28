/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    nodeMap := make(map[*Node]*Node) // key: original, val: copied
	anchorNode := &Node{}
	iter := anchorNode

	for head != nil {
		nodeMap[head] = &Node{
			Val: head.Val,
			Next: nil,
			Random: head.Random,
		}

		iter.Next = nodeMap[head]
		iter = iter.Next
		head = head.Next
	}

	// get randomness
	newIter := anchorNode.Next
	for newIter != nil {
		newIter.Random = nodeMap[newIter.Random]
		newIter = newIter.Next
	}


	return anchorNode.Next
}

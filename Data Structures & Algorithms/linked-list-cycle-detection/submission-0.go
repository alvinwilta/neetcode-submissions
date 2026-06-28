/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil{
		return false
	}
	if head.Val == 1001 {
		return true
	}

	head.Val = 1001
	return hasCycle(head.Next)
}

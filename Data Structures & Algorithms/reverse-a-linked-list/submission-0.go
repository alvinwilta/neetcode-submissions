/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	next := head.Next
	ret := reverseList(head.Next)
	next.Next = head
	head.Next = nil

	return ret
}

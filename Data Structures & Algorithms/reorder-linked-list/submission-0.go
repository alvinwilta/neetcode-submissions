/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}
    // slow and fast to find half
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	rev := reverseLinkedList(slow.Next)
	slow.Next = nil

	// merge 2 list
	for head != nil  {
		tmp1 := head.Next
		tmp2 := rev.Next
		head.Next = rev
		head = tmp1
		if head != nil {
			rev.Next = head
			rev = tmp2
		}
	}
}

func reverseLinkedList(head *ListNode) *ListNode {
    if head.Next == nil {
        return head
    }

    ret := reverseLinkedList(head.Next)
    head.Next.Next = head
    head.Next = nil

    return ret
}
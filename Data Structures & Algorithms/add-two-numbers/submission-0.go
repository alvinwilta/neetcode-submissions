/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ans := &ListNode{}
	iter := ans

	carry := false
	for !(l1 == nil && l2 == nil) {
		if l1 == nil {
			l1 = &ListNode{
				Val: 0,
			}
		}
		if l2 == nil {
			l2 = &ListNode{
				Val: 0,
			}
		}
		newVal := l1.Val + l2.Val
		if carry {
			newVal += 1
			carry = false
		}
		if newVal > 9 {
			newVal = newVal % 10
			carry =true
		}
		iter.Next = &ListNode{
			Val: newVal,
		}

		iter = iter.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if carry {
		iter.Next = &ListNode{
			Val: 1,
		}
	}

	return ans.Next
}

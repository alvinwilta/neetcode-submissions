/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	ret := head
    for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			head = head.Next
			list2 = list2.Next
			continue
		}
		if list2 == nil {
			head.Next = list1
			head = head.Next
			list1 = list1.Next
			continue
		}
		if list1.Val < list2.Val {
			head.Next = list1
			head = head.Next
			list1 = list1.Next
			continue
		} else {
			head.Next = list2
			head = head.Next
			list2 = list2.Next
			continue
		}
	}

	return ret.Next
}

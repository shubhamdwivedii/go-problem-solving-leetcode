/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var rev *ListNode
	rev = new(ListNode)

	for head != nil {
		rev.Val = head.Val
		newR := new(ListNode)
		newR.Next = rev
		rev = newR
		head = head.Next
	}

	return rev.Next
}
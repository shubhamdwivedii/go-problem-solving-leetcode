/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}

	var nth *ListNode
	var curr *ListNode
	nth = head
	curr = head

	dist := 0
	for curr != nil && dist < n-1 {
		curr = curr.Next
		dist++
	}

	for curr.Next != nil {
		curr = curr.Next
		nth = nth.Next
	}

	if head == nth {
		return nth.Next
	}

	curr = head
	for curr.Next != nth {
		curr = curr.Next
	}

	curr.Next = nth.Next

	return head
}
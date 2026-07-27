/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// TON SON
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

// Shubh 

func reverseList(head *ListNode) *ListNode {
    current := head 
    prev := current
    for current != nil {
        next := current.Next  
        if next == nil {
            break 
        }
        current.Next = next.Next 
        next.Next = prev
        prev = next 
    }

    return prev 
}

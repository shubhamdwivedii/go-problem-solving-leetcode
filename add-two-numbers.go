/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	output := new(ListNode)
	var opt *ListNode = &(*output)
	var lst1 *ListNode = &(*l1)
	var lst2 *ListNode = &(*l2)

	buff := 0
	for opt != nil {
		var n1, n2 int
		if lst1 != nil {
			n1 = lst1.Val
			lst1 = lst1.Next
		}

		if lst2 != nil {
			n2 = lst2.Val
			lst2 = lst2.Next
		}

		sum := n1 + n2 + buff
		if sum >= 10 {
			buff = sum / 10
			sum = sum % 10
		} else {
			buff = 0
		}
		opt.Val = sum
		if lst1 == nil && lst2 == nil && buff == 0 {
			opt.Next = nil
		} else {
			opt.Next = new(ListNode)
		}
		opt = opt.Next
	}

	return output
}
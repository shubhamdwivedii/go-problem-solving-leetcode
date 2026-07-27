/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	output := new(ListNode)
	opt := output

	for opt != nil {
		if l1 == nil && l2 == nil {
			opt = nil
			break
		}

		if l1 == nil {
			opt.Val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			opt.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val >= l2.Val {
			opt.Val = l2.Val
			l2 = l2.Next
		} else {
			opt.Val = l1.Val
			l1 = l1.Next
		}

		if l1 != nil || l2 != nil {
			opt.Next = new(ListNode)
		}

		opt = opt.Next
	}
	return output
}

// TOn SO1 


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := new(ListNode)
    current := dummy 

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            current.Next = list1 
            list1 = list1.Next
        } else {
            current.Next = list2 
            list2 = list2.Next
        }
        current = current.Next
    }

    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2 
    }

    return dummy.Next
}


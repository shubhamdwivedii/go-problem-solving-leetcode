/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	var hare *ListNode
	var tortoise *ListNode

	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	} else {
		hare = head.Next
		tortoise = head.Next
	}

	if hare.Next != nil {
		hare = hare.Next
	} else {
		return false
	}

	for hare != nil {
		if hare == tortoise {
			return true
		}

		hare = hare.Next
		if hare != nil && hare.Next != nil {
			hare = hare.Next
		}
		tortoise = tortoise.Next
	}

	return false
}

// TOn SO1
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    tortoise := head
    rabbit := head

    for rabbit != nil && rabbit.Next != nil {
        tortoise = tortoise.Next
        rabbit = rabbit.Next.Next

        if tortoise == rabbit {
            return true
        }
    }

    return false
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	encMap := make(map[*ListNode]bool)

	if headA == nil || headB == nil {
		return nil
	}

	for {
		if headA != nil {
			if _, ok := encMap[headA]; ok {
				return headA
			} else {
				encMap[headA] = true
			}
			headA = headA.Next
		}

		if headB != nil {
			if _, ok := encMap[headB]; ok {
				return headB
			} else {
				encMap[headB] = true
			}
			headB = headB.Next
		}

		if headA != nil && headB != nil {
			break
		}
	}
	return nil
}
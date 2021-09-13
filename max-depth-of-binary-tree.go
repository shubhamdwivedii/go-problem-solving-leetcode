/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return calcDepth(root, 1)
}

func calcDepth(node *TreeNode, curr int) int {
	if node.Left == nil && node.Right == nil {
		return curr
	}

	var leftDepth, rightDepth int
	if node.Left != nil {
		leftDepth = calcDepth(node.Left, curr+1)
	}
	if node.Right != nil {
		rightDepth = calcDepth(node.Right, curr+1)
	}

	if leftDepth >= rightDepth {
		return leftDepth
	}
	return rightDepth
}
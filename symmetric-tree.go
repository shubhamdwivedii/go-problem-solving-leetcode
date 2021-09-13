/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode != nil && rightNode != nil {
		if leftNode.Val == rightNode.Val {
			return isMirror(leftNode.Left, rightNode.Right) && isMirror(leftNode.Right, rightNode.Left)
		}
	}
	return false
}
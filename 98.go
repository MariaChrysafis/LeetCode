/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isOkay (root *TreeNode, left int, right int) bool {
    if root == nil {
        return true
    }
    if left >= root.Val || root.Val >= right {
        return false
    }
    return isOkay(root.Left, left, root.Val) && isOkay(root.Right, root.Val, right)
}
func isValidBST(root *TreeNode) bool {
    return isOkay(root, -2147483649, 2147483648)
}

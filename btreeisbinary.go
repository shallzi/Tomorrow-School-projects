package piscine

func BTreeIsBinary(root *TreeNode) bool {
	for root != nil {
		if root.Left.Data > root.Right.Data {
			return false
		} else {
			root = root.Left
		}
		root = root.Right
	}
	return true
}

package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == node {
		return rplc
	}

	if root.Data < node.Data {
		root.Right = BTreeTransplant(root.Right, node, rplc)
	}
	if root.Data > node.Data {
		root.Left = BTreeTransplant(root.Left, node, rplc)
	}
	return root
}

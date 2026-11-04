package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	var newNode *TreeNode
	if data < root.Data {
		newNode = BTreeInsertData(root.Left, data)
		root.Left = newNode
	} else {
		newNode = BTreeInsertData(root.Right, data)
		root.Right = newNode
	}
	newNode.Parent = root
	return root
}

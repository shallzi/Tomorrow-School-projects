package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		if root.Left == nil {
			newNode := BTreeInsertData(root.Left, data)
			root.Left = newNode
			newNode.Parent = root
		} else {
			return BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			newNode := BTreeInsertData(root.Right, data)
			root.Right = newNode
			newNode.Parent = root
		} else {
			return BTreeInsertData(root.Right, data)
		}
	}
	return root
}

package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return root
	}

	arr := []*TreeNode{root}

	for len(arr) > 0 {
		current := arr[0]
		arr = arr[1:]

		if current.Data == elem {
			return current
		}

		if current.Right != nil {
			arr = append(arr, current.Right)
		}
		if current.Left != nil {
			arr = append(arr, current.Left)
		}

	}
	return nil
}

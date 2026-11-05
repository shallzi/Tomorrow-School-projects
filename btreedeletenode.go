package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil || node == nil {
		return root
	}

	// Ищем нужный узел в дереве
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
		if root.Right != nil {
			root.Right.Parent = root
		}
	} else {
		// Узел найден

		// Случай 1: Нет детей или только один ребенок (слева)
		if root.Left == nil {
			temp := root.Right
			if temp != nil {
				temp.Parent = root.Parent
			}
			return temp
		}
		// Случай 2: Только один ребенок (справа)
		if root.Right == nil {
			temp := root.Left
			if temp != nil {
				temp.Parent = root.Parent
			}
			return temp
		}

		// Случай 3: Есть оба ребенка
		// Находим минимальный элемент в правом поддереве (successor)
		curr := root.Right
		for curr.Left != nil {
			curr = curr.Left
		}

		// Копируем данные найденного узла в текущий узел
		root.Data = curr.Data

		// Удаляем дублирующийся узел из правого поддерева
		root.Right = BTreeDeleteNode(root.Right, curr)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}

	return root
}

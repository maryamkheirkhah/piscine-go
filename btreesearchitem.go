package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	} else if root.Data == elem {
		return root
	} else {
		searchInLeft := BTreeSearchItem(root.Left, elem)
		if searchInLeft != nil {
			return searchInLeft
		}
		searchInRight := BTreeSearchItem(root.Right, elem)
		if searchInRight != nil {
			return searchInRight
		}
		return nil
	}
}

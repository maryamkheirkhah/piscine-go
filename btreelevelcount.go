package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	BTLC := BTreeLevelCount(root.Left)
	BTRC := BTreeLevelCount(root.Right)
	max := BTLC
	if BTRC > max {
		max = BTRC
	}
	return max + 1
}

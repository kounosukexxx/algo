package l776

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, target int) []*TreeNode {
	if root.Val <= target {
		if root.Right == nil {
			return []*TreeNode{root, nil}
		}

		bst := splitBST(root.Right, target)
		root.Right = bst[0]
		return []*TreeNode{root, bst[1]}
	}

	if root.Left == nil {
		return []*TreeNode{nil, root}
	}

	bst := splitBST(root.Left, target)
	root.Left = bst[1]
	return []*TreeNode{bst[0], root}
}

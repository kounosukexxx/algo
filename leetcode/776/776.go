package l776

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, target int) []*TreeNode {
	dummyRoot := &TreeNode{Val: target}
	recurse(dummyRoot, root, target)

	return []*TreeNode{dummyRoot.Left, dummyRoot.Right}
}

func recurse(dummyRoot, node *TreeNode, target int) {
	if node == nil {
		return
	}

	if node.Val <= target {
		right := node.Right
		node.Right = nil
		unite(dummyRoot, node)
		recurse(dummyRoot, right, target)
	} else {
		left := node.Left
		node.Left = nil
		unite(dummyRoot, node)
		recurse(dummyRoot, left, target)
	}
}

func unite(head, biasedTreeNode *TreeNode) {
	if biasedTreeNode.Val <= head.Val {
		if head.Left == nil {
			head.Left = biasedTreeNode
			return
		}
		unite(head.Left, biasedTreeNode)
	} else {
		if head.Right == nil {
			head.Right = biasedTreeNode
			return
		}
		unite(head.Right, biasedTreeNode)
	}
}

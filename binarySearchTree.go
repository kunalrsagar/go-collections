package collections

type BinarySearchTree[M Mix] struct {
	Root *TreeNode[M] `json:"root"`
}

func (bst *BinarySearchTree[M]) Insert(value M) {
	if bst.Root == nil {
		root := NewTreeNode(value)
		bst.Root = root
	} else {
		node := bst.Root
		CheckNInsert(value, node)
	}
}

func CheckNInsert[M Mix](value M, node *TreeNode[M]) {
	if value > node.Value {
		// move to right if value > Root.Value
		if node.Right == nil {
			// if Right node does not exit, create it
			newNode := NewTreeNode(value)
			node.Right = newNode
		} else {
			// check the value of node and insert on its left or right
			CheckNInsert(value, node.Right)
		}
	} else {
		// move to left if value < Root.Value
		if node.Left == nil {
			// if Left node does not exit, create it
			newNode := NewTreeNode(value)
			node.Left = newNode
		} else {
			// check the value of node and insert on its left or right
			CheckNInsert(value, node.Left)
		}
	}
}

func (bst *BinarySearchTree[M]) Lookup(value M) bool {
	node := bst.Root
	return checkNodeNValue(value, node)
}

func checkNodeNValue[M Mix](value M, node *TreeNode[M]) bool {
	if node != nil {
		if value > node.Value {
			// check on right node
			return checkNodeNValue(value, node.Right)
		} else if value < node.Value {
			// check on left node
			return checkNodeNValue(value, node.Left)
		} else {
			return true
		}
	} else {
		return false
	}
}

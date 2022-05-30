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
		for true {
			if value > node.Value {
				// move to right if value > Root.Value
				if node.Right == nil {
					// if Right node does not exit, create it
					newNode := NewTreeNode(value)
					node.Right = newNode
					return
				} else {
					// check the value of node and insert on its left or right
					node = node.Right
					continue
				}
			} else {
				// move to left if value < Root.Value
				if node.Left == nil {
					// if Left node does not exit, create it
					newNode := NewTreeNode(value)
					node.Left = newNode
					return
				} else {
					// check the value of node and insert on its left or right
					node = node.Left
					continue
				}
			}
		}
	}
}

func (bst *BinarySearchTree[M]) Lookup(value M) bool {
	node := bst.Root
	for node != nil {
		if value > node.Value {
			// check on right node
			node = node.Right
			continue
		} else if value < node.Value {
			// check on left node
			node = node.Left
			continue
		} else if node.Value == value {
			return true
		}
	}
	return false
}

package datastructures

type AVLNode struct {
	Val         interface{}
	Height      int
	Left, Right *AVLNode
}

func RotateLeft(n *AVLNode) *AVLNode {
	node := n.Right
	n.Right = node.Left
	node.Left = n
	n.Height = AVLMax(height(n.Left), height(n.Right)) + 1
	node.Height = AVLMax(height(node.Right), height(node.Left)) + 1
	return node
}

func RotateRight(n *AVLNode) *AVLNode {
	root := n.Left
	n.Left = root.Right
	root.Right = n

	root.Height = AVLMax(height(root.Left), height(root.Right)) + 1
	n.Height = AVLMax(height(n.Left), height(n.Right)) + 1

	return root
}

func RotateLeftRight(n *AVLNode) *AVLNode {
	n.Left = RotateLeft(n.Left)
	n = RotateRight(n)
	return n
}

func RotateRightLeft(n *AVLNode) *AVLNode {
	n.Right = RotateRight(n.Right)
	n = RotateLeft(n)
	return n
}

func height(node *AVLNode) int {
	if node != nil {
		return node.Height
	}
	return -1
}

func AVLMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Insert(n *AVLNode, key interface{}) *AVLNode {
	if n == nil {
		n = &AVLNode{key, 0, nil, nil}
		n.Height = AVLMax(height(n.Left), height(n.Right)) + 1
		return n
	}

	if key < n.Val {
		n.Left = Insert(n.Left, key)
		if height(n.Left)-height(n.Right) == 2 {
			if key < n.Left.Val {
				n = RotateRight(n)
			} else {
				n = RotateLeftRight(n)
			}
		}
	} else if key > n.Val {
		n.Right = Insert(n.Right, key)
		if height(n.Right)-height(n.Left) == 2 {
			if key > n.Right.Val {
				n = RotateLeft(n)
			} else {
				n = RotateRightLeft(n)
			}
		}
	}

	n.Height = AVLMax(height(n.Left), height(n.Right)) + 1
	return n
}

type action func(node *AVLNode)

func InOrder(root *AVLNode, action action) {
	if root == nil {
		return
	}

	InOrder(root.Left, action)
	action(root)
	InOrder(root.Right, action)
}

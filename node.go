package squarestore

import "fmt"

/*SimpleNode represents a simple node, linked to previous and next (if exists) elements into the structure.*/
type SimpleNode struct {
	next     *Node
	previous *Node
}

/*HorizontalNode represents the features of a Node in the main X axis, which will be linked to its left and right (if exists) nodes.*/
type HorizontalNode struct {
	left  *Node
	right *Node
}

/*VerticalNode represents the features of a Node in the main Y axis, which will be linked to its up (if exists) and down nodes.*/
type VerticalNode struct {
	up   *Node
	down *Node
}

/*DiagonalNode represents the features of a Node in the main diagonal, which will be linked to its previous and next (if exists) nodes in that main diagonal.*/
type DiagonalNode struct {
	dUp   *Node
	dDown *Node
}

/*Node represents a complete Node with all the posible link-features it can have.*/
type Node struct {
	SimpleNode
	HorizontalNode
	VerticalNode
	DiagonalNode
	active   bool
	position int64
	Content  interface{}
}

/*Assemble links the caller Node with its related Nodes.*/
func (n *Node) assemble(previous *Node, left *Node, down *Node, dDown *Node, c interface{}) error {
	if previous == nil {
		return fmt.Errorf("Previous node must not be nil")
	}

	n.previous = previous
	previous.next = n

	if left != nil {
		n.left = left
		left.right = n
	}

	if down != nil {
		n.down = down
		down.up = n
	}

	if dDown != nil {
		n.dDown = dDown
		dDown.dUp = n
	}

	n.active = true

	n.position = previous.position + 1

	n.Content = c

	return nil
}

func (n *Node) assembleFirst(c interface{}) {
	n.active = true

	n.position = 1

	n.Content = c
}

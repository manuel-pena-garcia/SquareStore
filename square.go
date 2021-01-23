package squarestore

import "fmt"

/*Square represents the SquareStore structure, pointing to its first and last Nodes*/
type Square struct {
	first *Node
	last  *Node
}

/*Insert allows to insert an element at the end of the structure*/
func (s *Square) Insert(c interface{}) {

	var n Node

	lastPosition := s.last.position + 1

	if s.first == nil {
		n.assembleFirst(c)
	} else {
		n.assemble(s.last, s.getLeftNode(lastPosition), s.getDownNode(lastPosition), s.getDDownNode(lastPosition), c)
	}
}

/*Read allows to read the element stored in a given position, except if it is deactivated, in which case nil will be returned*/
func (s *Square) Read(position int64) *Node {
	n := s.getNodeAt(position)

	if n == nil || n.active == false {
		return nil
	}

	return n
}

/*Activate allows to activate (logically undelete) the element stored in a given position*/
func (s *Square) Activate(position int64) {
	s.getNodeAt(position).active = true
}

/*Deactivate allows to deactivate (logically delete) the element stored in a given position*/
func (s *Square) Deactivate(position int64) {
	s.getNodeAt(position).active = false
}

func (s *Square) getNodeAt(position int64) *Node {
	fmt.Println(position)

	//TODO

	return nil
}

func (s *Square) getLeftNode(position int64) *Node {
	if position == 1 {
		return nil
	}

	if b, step := isInXAxis(position); b == true {
		return s.getNodeAt(step*step + 1)
	}

	return nil
}

func (s *Square) getDownNode(position int64) *Node {
	if position == 1 {
		return nil
	}

	if b, step := isInYAxis(position); b == true {
		return s.getNodeAt(step * step)
	}

	return nil
}

func (s *Square) getDDownNode(position int64) *Node {
	if position == 1 {
		return nil
	}

	if b, step := isInDiagonal(position); b == true {
		return s.getNodeAt(step*step - step + 1)
	}

	return nil
}

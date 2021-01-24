package squarestore

import "log"

/*Square represents the SquareStore structure, pointing to its first and last Nodes*/
type Square struct {
	first   *Node
	last    *Node
	maxroot int64
}

/*Insert allows to insert an element at the end of the structure*/
func (s *Square) Insert(c interface{}) {

	var n Node

	if s.first == nil {
		n.assembleFirst(c)

		s.first = &n
		s.last = &n
	} else {
		lastPosition := s.last.position + 1

		n.assemble(s.last, s.getLeftNode(lastPosition), s.getDownNode(lastPosition), s.getDDownNode(lastPosition), c)

		s.last = &n
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
	origin, path, steps, err := s.retrieveBestPath(position)

	if err != nil {
		log.Fatal(err)

		return nil
	}

	var n *Node

	stepSq := steps * steps
	stepSqD := steps*steps - steps + 1

	if origin == beginning {
		n = s.first

		if path == xAxis {
			for i := int64(0); i < steps; i++ {
				n = n.right
			}
			for i := stepSq + 1; i < position; i++ {
				n = n.next
			}
		} else if path == yAxis {
			for i := int64(0); i < steps; i++ {
				n = n.up
			}
			for i := stepSq; i > position; i-- {
				n = n.previous
			}
		} else {
			for i := int64(0); i < steps; i++ {
				n = n.dUp
			}
			if position > stepSqD {
				for i := stepSqD; i < position; i++ {
					n = n.next
				}
			} else if position < stepSqD {
				for i := stepSqD; i > position; i-- {
					n = n.previous
				}
			}
		}
	} else {
		//TODO -> new Path: LAST??

		n = s.last

		if path == xAxis {

		} else if path == yAxis {

		} else {

		}
	}

	return n
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
		s.maxroot = step + 1

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

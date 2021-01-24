package squarestore

import (
	"fmt"
	"math"
)

/*Path represents the option select between X and Y axis or main diagonal to go through the SquareStore structure.*/
type Path uint

const (
	xAxis Path = iota
	yAxis
	diagonal
)

/*Origin represents the starting point of a way through the SquareStore structure.*/
type Origin uint

const (
	beginning Origin = iota
	end
)

/*RetrieveBestPath retrieves the fastest way to get to a given position through the SquareStore structure*/
func (s *Square) retrieveBestPath(position int64) (Origin, Path, int64, error) {
	var err error

	if s.last == nil {
		err = fmt.Errorf("The SquareStore is empty")
	} else if position > s.last.position {
		err = fmt.Errorf("Index out of bounds")
	}

	root := int64(math.Sqrt(float64(position)))

	rest := position - root*root

	var first Origin
	var second Path

	if s.maxroot-root > root {
		first = beginning
	} else {
		first = end
	}

	if rest == 0 {
		second = yAxis
	} else if rest == 1 {
		second = xAxis
	} else if rest > (root/2 + 1) {
		if (root+1)*(root+1) <= s.last.position && rest > 3*root/2 {
			second = yAxis
		} else {
			second = diagonal
		}
	} else {
		second = xAxis
	}

	return first, second, root, err
}

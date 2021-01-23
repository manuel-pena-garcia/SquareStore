package squarestore

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

func isInXAxis(position int64) (bool, int64) {
	x, y := isPerfectSquare(position - 1)

	return x, y + 1
}

func isInYAxis(position int64) (bool, int64) {
	return isPerfectSquare(position)
}

func isInDiagonal(position int64) (bool, int64) {
	if position == 1 {
		return true, 0
	}

	return checkInDiagonal(position, 2, 1)
}

func checkInDiagonal(n int64, t int64, p int64) (bool, int64) {
	/*The expression a number should match to belong to the main diagonal.*/
	rs := t*t - t + 1

	if n > rs {
		/*n is still greater than the value to match, we must iterate again.*/
		return checkInDiagonal(n, t*t, t)
	} else if n < rs {
		/*The result is bounded between p (which is equal to sqrt(t)) and t*/
		return checkInDiagonalBetweenBounds(n, p, t)
	} else {
		/*We found that n == rs, it matches the given expression*/
		return true, t - 1
	}
}

func checkInDiagonalBetweenBounds(n int64, left int64, right int64) (bool, int64) {
	if left >= right {
		return false, (2*left - right)
	}

	t := left + (right-left)/2

	if t == left {
		return right*right == n, t - 1
	}

	/*The expression a number should match to belong to the main diagonal.*/
	rs := t*t - t + 1

	if n > rs {
		/*The potential result, if exists, belongs to the second half*/
		return checkInDiagonalBetweenBounds(n, t, right)
	} else if n < rs {
		/*The potential result, if exists, belongs to the first half*/
		return checkInDiagonalBetweenBounds(n, left, t)
	} else {
		/*We have a match as n == rs*/
		return true, t - 1
	}
}

func RetrieveBestPath(position int64) (Origin, Path) {
	//TODO
	return 0, 0
}

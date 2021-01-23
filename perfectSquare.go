package squarestore

import "fmt"

/*isPerfectSquare checks if a given number n is a perfect square.*/
func isPerfectSquare(n int64) (bool, int64) {
	if n == 1 {
		return true, 0
	}

	return checkPerfectSquare(n, 2, 1)
}

func checkPerfectSquare(n int64, t int64, p int64) (bool, int64) {
	sq := t * t

	if n > sq {
		/*n is still greater than sq, so we must continue the iteration.*/
		return checkPerfectSquare(n, sq, t)
	} else if n < sq {
		/*We have bounded the potential root between p (which is equal to sqrt(t)) and t.*/
		return checkBetweenBounds(n, p, t)
	} else {
		/*We have found that n = t * t, so we return true and the number of steps to achieve position n.*/
		return true, t - 1
	}
}

func checkBetweenBounds(n int64, left int64, right int64) (bool, int64) {
	fmt.Println("LEFT:", left, ", RIGHT:", right)

	if left >= right {
		/*We determined n is not a perfect square, so we return false.*/
		return false, (2*left - right)
	}

	t := left + (right-left)/2

	if t == left {
		return right*right == n, t - 1
	}

	sq := t * t

	if n > sq {
		/*The potential root will be in the right half.*/
		return checkBetweenBounds(n, t, right)
	} else if n < sq {
		/*The potential root will be in the left half.*/
		return checkBetweenBounds(n, left, t)
	} else {
		/*We have found that n = t * t, so we return true.*/
		return true, t - 1
	}
}

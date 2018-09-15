package diffsquares

func Difference(n int) (diff int) {
	var sumOfSq, sqOfSum int
	for i := 1; i <= n; i++ {
		sumOfSq += i * i
		sqOfSum += i
	}
	return sqOfSum*sqOfSum - sumOfSq
}

func SquareOfSums(n int) (sqOfSum int) {
	for i := 1; i <= n; i++ {
		sqOfSum += i
	}
	return sqOfSum * sqOfSum
}

func SumOfSquares(n int) (sumOfSq int) {
	for i := 1; i <= n; i++ {
		sumOfSq += i * i
	}
	return sumOfSq
}

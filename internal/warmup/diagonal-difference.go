package warmup

//Given a square matrix, calculate the absolute difference between the sums of its diagonals.
//
//The left-to-right diagonal = . The right to left diagonal = . Their absolute difference is .
//
//https://www.hackerrank.com/challenges/diagonal-difference/problem

func DiagonalDifference(arr [][]int32) int32 {

	var firstDiagonal, secondDiagonal int32

	for i := range arr {
		for j := range arr {
			if i == j {
				firstDiagonal += arr[i][j]
			}
			if i+j == (len(arr) - 1) {
				secondDiagonal += arr[i][j]
			}
		}
	}

	if firstDiagonal < secondDiagonal {
		return secondDiagonal - firstDiagonal
	}
	return firstDiagonal - secondDiagonal
}

package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [3][3]int32) int32 {
	var primaryDiagonal, secondaryDiagonal float64

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == j {
				primaryDiagonal += float64(arr[i][j])
			}
			if (i + j) == (len(arr) - 1) {
				secondaryDiagonal += float64(arr[i][j])
			}
		}
	}

	difference := math.Abs(primaryDiagonal - secondaryDiagonal)
	return int32(difference)
}

func main() {
	var arr = [3][3]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	result := diagonalDifference((arr))

	fmt.Println(result)
}

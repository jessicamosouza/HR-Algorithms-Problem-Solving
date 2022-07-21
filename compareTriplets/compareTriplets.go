package main

import (
	"fmt"
)

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a, b []int32) []int32 {
	var final_points = []int32{0, 0}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			final_points[0]++
		} else if a[i] < b[i] {
			final_points[1]++
		}
	}

	return final_points
}

func main() {
	a := []int32{1, 2, 3}
	b := []int32{3, 2, 1}
	results := compareTriplets(a, b)
	fmt.Println(results)
}


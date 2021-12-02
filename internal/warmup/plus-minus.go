package warmup

import (
	"fmt"
)

//Given an array of integers, calculate the ratios of its elements that
//are positive, negative, and zero. Print the decimal value of each
//fraction on a new line with  places after the decimal.

//Note: This challenge introduces precision problems. The test cases
//are scaled to six decimal places, though answers with absolute error
//of up to  are acceptable.

//https://www.hackerrank.com/challenges/plus-minus/problem

func PlusMinus(arr []int32) {

	var positives, zeros, negatives int32
	size := float64(len(arr))

	for i := range arr {
		if arr[i] > 0 {
			positives++
			continue
		}
		if arr[i] == 0 {
			zeros++
		} else {
			negatives++
		}
	}

	positivesRate := float64(positives) / size
	zerosRate := float64(zeros) / size
	negativesRate := float64(negatives) / size

	fmt.Printf("%.6f \n", positivesRate)
	fmt.Printf("%.6f \n", negativesRate)
	fmt.Printf("%.6f \n", zerosRate)

}

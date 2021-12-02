package warmup

import "fmt"

//Given five positive integers, find the minimum and maximum values
//that can be calculated by summing exactly four of the five integers.
//Then print the respective minimum and maximum values as a single
//line of two space-separated long integers.

func MiniMaxSum(arr []int64) {

	min := arr[0]
	max := arr[0]

	sum := int64(0)
	for _, v := range arr {
		sum += v
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	fmt.Println(fmt.Sprintf("%d %d", sum-max, sum-min))
}

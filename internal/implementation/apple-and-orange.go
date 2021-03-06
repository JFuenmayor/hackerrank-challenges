package implementation

import (
	"fmt"
	"log"
)

//Sam's house has an apple tree and an orange tree that yield an
//abundance of fruit. Using the information given below, determine
//the number of apples and oranges that land on Sam's house.
//In the diagram below:

//1)The red region denotes the house, where *s* is the start point, and
//*t* is the endpoint. The apple tree is to the left of the house, and
//the orange tree is to its right.
//2)Assume the trees are located on a single point, where the apple
//tree is at point *a*, and the orange tree is at point *b*.
//3)When a fruit falls from its tree, it lands *d* units of distance from its
//tree of origin along the x-axis. *A negative value of *d* means the
//fruit fell *d* units to the tree's left, and a positive value of *d* means
//it falls *d*  units to the tree's right. *

//Given the value of  for  apples and  oranges, determine how
//many apples and oranges will fall on Sam's house (i.e., in the
//inclusive range [s,t])?

//https://www.hackerrank.com/challenges/apple-and-orange/problem

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applesOnRange, orangesOnRange int

	for _, apple := range apples {
		position := a + apple
		if position <= t && position >= s {
			applesOnRange++
		}
	}

	for _, orange := range oranges {
		position := b + orange
		if position <= t && position >= s {
			orangesOnRange++
		}
	}

	log.Println(fmt.Sprintf("%d\n%d", applesOnRange, orangesOnRange))
}

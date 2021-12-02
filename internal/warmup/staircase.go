package warmup

import (
	"fmt"
	"log"
	"strings"
)

//Its base and height are both equal to . It is drawn using # symbols
//and spaces. The last line is not preceded by any spaces.
//Write a program that prints a staircase of size .

//Function Description
//Complete the staircase function in the editor below.
//staircase has the following parameter(s):
//int n: an integer

//Print a staircase as described above.

func Staircase(n int32) {
	for i := int32(0); i < n; i++ {
		empty := strings.Repeat(" ", int(n-i)-1)
		hash := strings.Repeat("#", int(i)+1)
		stairs := fmt.Sprintf("%s%s", empty, hash)
		log.Println(stairs)
	}
}

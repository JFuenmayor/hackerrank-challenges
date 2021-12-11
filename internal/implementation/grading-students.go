package implementation

import (
	"log"
	"math"
)

//HackerLand University has the following grading policy:

//1) Every student receives a grade in the inclusive range from 0 to
//100.
//2) Any  less than 40 is a failing grade.

//Sam is a professor at the university and likes to round each student's
//grade according to these rules:

//1) If the difference between the grade and the next multiple of 5 is
//less than 3, round grade up to the next multiple of 5.
//2)If the value of grade is less than 38, no rounding occurs as the
//result will still be a failing grade.

func GradingStudents(grades []int32) []int32 {

	for i, v := range grades {
		if v < 38 {
			grades[i] = v
			continue
		}
		t := int32(math.Ceil(float64(v)/5)) * 5
		if t-v >= 3 {
			grades[i] = v
			continue
		}
		grades[i] = t
	}
	log.Println(grades)
	return grades
}

// best found https://www.hackerrank.com/rest/contests/master/challenges/grading/hackers/cmouli_84/download_solution?primary=true

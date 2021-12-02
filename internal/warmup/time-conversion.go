package warmup

import (
	"fmt"
	"time"
)

//Given a time in 12-hour AM/PM format, convert it to military (24-hour) time.
//
//Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
//- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

//https://www.hackerrank.com/challenges/time-conversion/problem

func TimeConversion (s string) string {
	layout := "03:04:05PM"
	result,_ := time.Parse(layout,s)

	return fmt.Sprintf("%s",result.Format("15:04:05"))
}

package warmup

//You are in charge of the cake for a child's birthday. You have decided
//the cake will have one candle for each year of their total age. They
//will only be able to blow out the tallest of the candles.
//Count how many candles are tallest.

//https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func BirthdayCakeCandles(candles []int32) int32 {

	// Modified from leaderboard

	max := candles[0]
	frequency := int32(0)
	for _, v := range candles {
		if v > max{
			max = v
			frequency = 0
		}
		if v == max {
			frequency++
		}
	}

	return frequency
}

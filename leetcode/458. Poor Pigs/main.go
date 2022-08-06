package main

import "math"

func main() {
	println(poorPigs(1000, 15, 60))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	pigs := 0
	for math.Pow(float64(minutesToTest/minutesToDie+1), float64(pigs)) < float64(buckets) {
		pigs++
	}

	return pigs
}
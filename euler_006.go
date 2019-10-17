package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(dif(100))
}

func sumofsquare(n int) int {
	var sq float64
	sum := 0
	for i := 1; i <= n; i++ {
		sq = math.Pow(float64(i), 2)
		var sb int = int(sq)
		sum += sb
	}
	return sum
}

func squareofsum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	var sum1 float64 = float64(sum)
	var sqsum float64 = math.Pow(sum1, 2)
	return int(sqsum)
}

func dif(a int) int {
	return squareofsum(a) - sumofsquare(a)
}

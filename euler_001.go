package main

import "fmt"

func main() {
	fmt.Println(ThreeOrFiveSum(1000))
}

//GaussSum to sum wih Gauss
func GaussSum(n int) int {
	return ((n + 1) * n) / 2
}

//ThreeOrFiveSum is command to sum
func ThreeOrFiveSum(x int) int {
	x--
	SumThree := GaussSum(x/3) * 3
	SumFive := GaussSum(x/5) * 5
	SumFifteen := GaussSum(x/15) * 15
	return SumThree + SumFive - SumFifteen
}

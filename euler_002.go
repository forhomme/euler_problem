package main

import "fmt"

func main() {
	fmt.Println(EvenFibSum(1000))
}

//EvenFibSum do initial
func EvenFibSum(limit int64) int64 {
	return FibSum(limit, 0, 1, 0)
}

//FibSum check even number
func FibSum(limit, prev, curr, EvenSum int64) int64 {
	if limit == 0 {
		return EvenSum
	} else if isEven(curr) {
		EvenSum += curr
	}
	return FibSum(limit-1, curr, prev+curr, EvenSum)
}

func isEven(n int64) bool {
	return n%2 == 0
}

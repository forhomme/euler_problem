package main

import (
	"fmt"
)

func main() {
	fmt.Println(primenumber(10001))
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primenumber(n int) int {
	prm := 0
	i := 2
	for {
		if prm >= n {
			break
		} else if isPrime(i) {
			prm++
		}
		i++
	}
	return i - 1
}

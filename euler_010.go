package main

import "fmt"

func main() {
	fmt.Println(sumPrime(10))
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumPrime(n int) int {
	sum := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

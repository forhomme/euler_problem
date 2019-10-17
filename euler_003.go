package main

import "fmt"

func main() {
	fmt.Println(largestprimefactor(100))
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func largestprimefactor(n int) int {
	max := 0
	for i := 2; i < n; i++ {
		if n%i == 0 {
			if isPrime(i) {
				max = i
			}
		}
	}
	return max
}

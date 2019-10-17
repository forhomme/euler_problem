package main

import "fmt"

func main() {
	fmt.Println(smallpositivenumber(100000000))
}

func isDivide(n int) bool {
	div := 0
	for i := 1; i <= 15; i++ {
		if n%i == 0 {
			div++
		}
	}
	if div == 15 {
		return true
	} else {
		return false
	}
}

func smallpositivenumber(n int) int {
	pos := 0
	for i := 2520; i < n; i++ {
		if isDivide(i) {
			pos = i
			break
		}
	}
	return pos
}

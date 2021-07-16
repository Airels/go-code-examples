package main

import (
	"fmt"
)

func isPrimeNumber(n int) bool {
	for i := n-1; i > 1; i-- {
		if n%i == 0 {
			return false;
		}
	}

	return true;
}

func findPrimeNumbers(n int, m int, finished chan bool) {
	for i := n; i <= m; i++ {
		if isPrimeNumber(i) {
			fmt.Println(i, "is a prime number")
		}
	}

	finished <- true
}

func main() {
	finished := make(chan bool)
	finished2 := make(chan bool)
	finished3 := make(chan bool)

	go findPrimeNumbers(10000000, 10000100, finished)
	go findPrimeNumbers(10000100, 10000200, finished2)
	go findPrimeNumbers(10000200, 10000300, finished3)

	<- finished
	<- finished2
	<- finished3
	fmt.Println("Done.")
}

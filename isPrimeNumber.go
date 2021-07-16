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

func main() {
	fmt.Println(isPrimeNumber(17))
}

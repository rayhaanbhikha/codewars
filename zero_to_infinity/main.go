package main

import "fmt"

func main() {
	Going(6)
}

func Going(n int) float64 {

	currentSum, currentBase := 1.0, 1.0

	for i := 0.0; i < n-1; i++ {
		currentBase *= float64(n - i)
		currentSum += (1 / currentBase)
	}
	v := int(currentSum * 10e5)
	fmt.Println(v)
	return float64(v) / 10e5
}

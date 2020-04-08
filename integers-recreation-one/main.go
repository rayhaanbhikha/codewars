package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	// fmt.Println(ListSquared(1, 42))
	// fmt.Println(ListSquared(42, 250))
	// fmt.Println(divisors(43))
	fmt.Println(divisors(84))
	// fmt.Println(divisors(3455))
	// fmt.Println(divisors(2345))
}

type pair []int

func ListSquared(m, n int) [][]int {
	pairs := make([][]int, 0)
	for i := m; i <= n; i++ {
		if pair, ok := divisors(i); ok {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func divisors(n int) (pair, bool) {
	sum := float64(math.Pow(float64(n), 2))
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			fmt.Println(i, " ")
			sum += math.Pow(float64(i), 2)
		}
	}
	sqrt := math.Sqrt(sum)
	if sqrt == math.Trunc(sqrt) {
		return pair{n, int(sum)}, true
	}
	return pair{}, false
}

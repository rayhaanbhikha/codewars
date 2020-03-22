package main

func SumCubes(n int) int {
	// your code here
	sum := 0
	for i := 1; i <= n; i++ {
		sum += cube(i)
	}

	return sum
}

func cube(x int) int {
	return x * x * x
}

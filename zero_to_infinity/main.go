package main

func main() {
	Going(5)
}

func Going(n int) float64 {

	var currentSum float64 = 1
	var currentBase float64 = 1

	for i := 0; i < n-1; i++ {
		newN := float64(n - i)
		currentBase *= newN
		appendVal := divide(1.0, currentBase)
		currentSum += appendVal
	}

	return currentSum
}

func divide(x, y float64) float64 {
	return x / y
}

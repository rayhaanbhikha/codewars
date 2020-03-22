package main

func main() {
	Going(6)
}

func Going(n int) float64 {

	currentSum, currentBase := 1.0, 1.0

	for i := 0; i < n-1; i++ {
		newN := float64(n - i)
		currentBase *= newN
		appendVal := 1 / currentBase
		currentSum += appendVal
	}
	v := int(currentSum * 10e5)
	return float64(v) / 10e5
}

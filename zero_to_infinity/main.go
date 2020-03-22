package main

func main() {
	Going(6)
}

func Going(n int) float64 {
	currentSum, currentBase := 1.0, 1.0
	for i := float64(n); i > 1.0; i-- {
		currentBase *= i
		currentSum += (1 / currentBase)
	}
	v := int(currentSum * 10e5)
	return float64(v) / 10e5
}

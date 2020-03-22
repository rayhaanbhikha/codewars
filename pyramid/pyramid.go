package main

func Pyramid(n int) [][]int {
	pyramid := make([][]int, n)
	for i := range pyramid {
		for j := 0; j <= i; j++ {
			pyramid[i] = append(pyramid[i], 1)
		}
	}
	return pyramid
}

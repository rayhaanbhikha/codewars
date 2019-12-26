package main

import "fmt"

func main() {
	// Decompose("21/23")
	fractions := []fraction{
		fraction{1, 2}, fraction{1, 3}, fraction{1, 13}, fraction{1, 598}, fraction{1, 897},
	}

	// fmt.Println(fractions)
	// computedFraction := addFraction(fractions[0], fractions[1])
	// newFractions := append([]fraction{computedFraction}, fractions[2:]...)
	// fmt.Println(computedFraction)
	// fmt.Println(newFractions)
	// fmt.Println(addFraction(fractions[3], fractions[4]))
	// fmt.Println(computeGCF(38201436, 41839668))
	fmt.Println(addNFractions(fractions))

	// Decompose("0.66")
	// Decompose("0.5")
	// Decompose("0.578")
	// Decompose("1")
	// Decompose("5")
	// Decompose("0")
}

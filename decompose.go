package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parse(s string) fraction {

	if strings.Contains(s, "/") {
		v := strings.Split(s, "/")
		num, _ := strconv.Atoi(v[0])
		denom, _ := strconv.Atoi(v[1])
		return NewFraction(num, denom)
	}

	number, _ := strconv.Atoi(s)
	switch {
	case number < 1:
		v := strings.Split(s, ".")
		num, _ := strconv.Atoi(v[1])
		denom := math.Pow(10, float64(len(v[1])))
		return NewFraction(num, int(denom))
	case number == 1:
		return NewFraction(1, 1)
	default:
		return NewFraction(number, 1)
	}
}

type fraction struct {
	num, denom int
}

func NewFraction(num, denom int) fraction {
	gcf := computeGCF(num, denom)
	return fraction{num / gcf, denom / gcf}
}

func (f fraction) isImproper() bool {
	return f.num > f.denom
}

func Decompose(s string) []string {
	if s == "0" {
		return []string{}
	}
	parsedFraction := parse(s)
	fmt.Println(parsedFraction)

	denominators := make([]fraction, 0)

	// improper fraction
	if parsedFraction.isImproper() {
		denominators = append(denominators, NewFraction(1, 1))
	}

	return []string{}
}
func addNFractions(fractions []fraction) fraction {
	n := len(fractions) - 1
	fmt.Println(fractions, n)

	if n >= 2 {
		computedFraction := addFraction(fractions[0], fractions[1])
		newFractions := append([]fraction{computedFraction}, fractions[2:]...)
		return addNFractions(newFractions)
	}

	if n == 1 {
		return addFraction(fractions[0], fractions[1])
	}

	return fractions[0]
}

func addFraction(f1, f2 fraction) fraction {
	num := (f1.num * f2.denom) + (f2.num * f1.denom)
	denom := f1.denom * f2.denom

	return NewFraction(num, denom)
}

func computeGCF(num1, num2 int) int {
	num1Factors := make([]int, 0)
	num2Factors := make([]int, 0)

	for i := num1; i != 0; i-- {
		if num1%i == 0 {
			num1Factors = append(num1Factors, i)
		}
	}
	for i := num2; i != 0; i-- {
		if num2%i == 0 {
			num2Factors = append(num2Factors, i)
		}
	}

	for _, v1 := range num1Factors {
		for _, v2 := range num2Factors {
			if v1 == v2 {
				return v1
			}
		}
	}
	return 1
}

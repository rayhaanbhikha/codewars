package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func IsValidCoordinates(coordinates string) bool {
	newCoordinates := strings.Split(coordinates, ", ")

	if len(newCoordinates) == 2 && validCoord(newCoordinates[0], 0, 90) && validCoord(newCoordinates[1], 0, 180) {
		return true
	}

	return false
}

func validCoord(lattitude string, lowerThres, upperThres float64) bool {

	if strings.Contains(lattitude, "e") {
		return false
	}

	v, err := strconv.ParseFloat(lattitude, 32)
	if err != nil {
		return false
	}

	if newV := math.Abs(v); newV > lowerThres && newV < upperThres {
		fmt.Println(v)
		return true
	}

	return false
}

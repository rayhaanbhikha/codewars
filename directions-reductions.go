package main

import "fmt"

func main() {

	input := []string{
		"NORTH",
		"SOUTH",
		"SOUTH",
		"EAST",
		"WEST",
		"NORTH",
		"WEST",
	}

	fmt.Println(isOpposite("NORTH", "SOUTH"))
	fmt.Println(isOpposite("SOUTH", "SOUTH"))
	fmt.Print(DirReduc(input))
}

var cancel = map[string]string{
	"NORTH": "SOUTH",
	"SOUTH": "NORTH",
	"EAST":  "WEST",
	"WEST":  "EAST",
}

func DirReduc(arr []string) []string {
	reducedDir := []string{arr[0]}
	for i := 1; i < len(arr); i++ {
		dir := arr[i]
		rLen := len(reducedDir)
		if rLen > 0 && isOpposite(dir, reducedDir[rLen-1]) {
			reducedDir = reducedDir[:rLen-1]
			continue
		}
		reducedDir = append(reducedDir, dir)
	}
	return reducedDir
}

func isOpposite(cDir, newDir string) bool {
	v, ok := cancel[cDir]
	return ok && v == newDir
}

package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
}

func getPath(x, y int, path *[]Point) bool {
	if x == 0 && y == 0 {
		return true
	}
	var success bool
	if isFree(x-1, y) {
		success = getPath(x-1, y, path)
	}
	if !success && isFree(x, y-1) {
		success = getPath(x, y-1, path)
	}
	if success {
		*path = append(*path, Point{x, y})
	}
	return success
}

func isFree(x, y int) bool {
	if x < 0 || y < 0 {
		return false
	} else {
		return grid[y][x] == 0
	}
}

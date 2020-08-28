package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("day3-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	grid := make(map[Coord]bool)
	coord := Coord{0, 0}

	for _, direction := range strings.Split(string(data), "") {
		switch direction {
		case ">":
			coord = Coord{coord.x + 1, coord.y}
		case "<":
			coord = Coord{coord.x - 1, coord.y}
		case "^":
			coord = Coord{coord.x, coord.y + 1}
		case "v":
			coord = Coord{coord.x, coord.y - 1}
		}
		grid[coord] = true
	}
	fmt.Println(len(grid))
}

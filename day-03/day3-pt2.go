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
	rcoord := Coord{0, 0}
	for index, direction := range strings.Split(string(data), "") {
		if index%2 == 0 {
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
		}
		if index%2 == 1 {
			switch direction {
			case ">":
				rcoord = Coord{rcoord.x + 1, rcoord.y}
			case "<":
				rcoord = Coord{rcoord.x - 1, rcoord.y}
			case "^":
				rcoord = Coord{rcoord.x, rcoord.y + 1}
			case "v":
				rcoord = Coord{rcoord.x, rcoord.y - 1}
			}
		}
		grid[coord] = true
		grid[rcoord] = true
	}
	fmt.Println(len(grid))
}

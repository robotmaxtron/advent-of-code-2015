package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day2-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	presents := strings.Split(string(data), "\n")

	total := 0

	for _, present := range presents {
		sides := []int{}

		for _, sideArea := range strings.Split(present, "x") {
			side, _ := strconv.Atoi(sideArea)
			sides = append(sides, side)
		}

		sort.Ints(sides)

		total += sides[0] * sides[1] * 3
		total += sides[0] * sides[2] * 2
		total += sides[1] * sides[2] * 2
	}

	fmt.Println(total)
}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day5-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	for _, needseval := range strings.Split(string(data), "") {
		// Must not contain  'ab' 'cd' 'pq' 'xy'
		if needseval = regexp.MustCompile("ab|cd|pq|xy"){
			return
		}
		// Must contain 3 vowels (aeiou)
		if needseval =
		// Must contain a double letter
}
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day-05-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	strings := strings.Split(string(data), "\n")

	count := 0

	for _, str := range strings {

		if validateString(str) {
			count++
		}

	}

	fmt.Print(count)
}

func validateString(str string) bool {
	return twoPairs(str) && repeatingLetterPattern(str)
}

func twoPairs(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if strings.Count(str, str[i:i+2]) >= 2 {
			return true
		}
	}
	return false
}

func repeatingLetterPattern(str string) bool {
	for pete := 0; pete < len(str)-2; pete++ {
		if str[pete] == str[pete+2] {
			return true
		}
	}
	return false
}

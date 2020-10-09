package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	return threeVowels(str) && doubleLetters(str) && banList(str)
}

func threeVowels(str string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	count := 0

	for _, vowel := range vowels {
		count += strings.Count(str, vowel)
	}

	if count >= 3 {
		return true
	}
	return false
}

func doubleLetters(str string) bool {
	for pete := 0; pete < len(str)-1; pete++ {
		if str[pete] == str[pete+1] {
			return true
		}
	}
	return false
}

func banList(str string) bool {
	return !regexp.MustCompile("ab|cd|pq|xy").MatchString(str)
}

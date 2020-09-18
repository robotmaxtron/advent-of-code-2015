package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
)

func main() {
	data, err := ioutil.ReadFile("day4-input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	number := 0
	regex := regexp.MustCompile(`\A000000`)
	digest := ""

	for !regex.MatchString(digest) {
		number++
		hash := md5.New()
		io.WriteString(hash, string(data))
		io.WriteString(hash, fmt.Sprintf("%d", number))
		digest = fmt.Sprintf("%x", hash.Sum(nil))
	}
	fmt.Println(number)
}

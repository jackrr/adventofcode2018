package main

import (
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	inputString := string(input)
	var threeLetter, twoLetter int

	letters := make(map[rune]int)
	for _, c := range inputString {
		if c == '\n' {
			// check for double
			for _, count := range letters {
				if count == 2 {
					twoLetter++
					break
				}
			}

			// triple
			for _, count := range letters {
				if count == 3 {
					threeLetter++
					break
				}
			}

			letters = make(map[rune]int)
			continue
		}

		letters[c]++
	}

	log.Printf("Found %d two-letters, %d three-letters, checksum: %d", twoLetter, threeLetter, twoLetter*threeLetter)
}

package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func reactive(a rune, b rune) bool {
	if a < b {
		return a+32 == b
	} else {
		return b+32 == a
	}
}

func remove(str string, start, length int) string {
	if start == 0 {
		return str[length:len(str)]
	}
	if start+length == len(str) {
		return str[0:start]
	}
	var sb strings.Builder
	sb.WriteString(str[0:start])
	sb.WriteString(str[start+length : len(str)])
	return sb.String()
}

func removepair(str string) (string, bool) {
	var prevIdx int
	var prevChar rune

	for idx, char := range str {
		if idx == 0 {
			prevChar = char
			prevIdx = idx
			continue
		}

		if reactive(prevChar, char) {
			return remove(str, prevIdx, 2), true
		}

		prevChar = char
		prevIdx = idx
	}

	return str, false
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	polymer := string(input)
	start := len(polymer)
	destroyed := true

	for destroyed {
		polymer, destroyed = removepair(polymer)
	}

	log.Printf("was: %d, now: %d", start, len(polymer))
}

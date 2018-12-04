package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func compare(id, oth []rune) (string, bool) {
	// log.Printf("Comparing %s and %s", string(id), string(oth))
	violations := 0
	var matched strings.Builder
	for idx, c := range id {
		if oth[idx] != c {
			violations++
		} else {
			matched.WriteRune(c)
		}
	}

	return matched.String(), violations < 2
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	id := make([]rune, 26)
	ids := make([][]rune, 0)
	idx := 0
	for _, c := range string(input) {
		if c == '\n' {
			// check match
			for idx, other := range ids {
				if overlap, match := compare(id, other); match {
					log.Printf("Matched %s and %s", string(id), string(other))
					log.Printf("Match found between ids #%d and #%d. Overlap: %s", idx, len(ids), overlap)
				}
			}

			// prepare for next id
			ids = append(ids, id)
			id = make([]rune, 26)
			idx = 0
			continue
		}

		id[idx] = c
		idx++
	}
}

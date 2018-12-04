package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	inputString := string(input)
	boxes := strings.Split(inputString, "\n")

	grid := make([][]int, 2000)
	for i := 0; i < 2000; i++ {
		grid[i] = make([]int, 2000)
	}

	var id, x, y, w, h int
	pureIds := make(map[int]bool, 0)

	for _, box := range boxes {
		tokens := strings.Split(box, " ")
		id, err = strconv.Atoi(strings.Replace(tokens[0], "#", "", 1))
		x, err = strconv.Atoi(strings.Split(tokens[2], ",")[0])
		y, err = strconv.Atoi(strings.Replace(strings.Split(tokens[2], ",")[1], ":", "", 1))
		if err != nil {
			panic(err)
		}

		w, err = strconv.Atoi(strings.Split(tokens[3], "x")[0])
		h, err = strconv.Atoi(strings.Split(tokens[3], "x")[1])
		if err != nil {
			panic(err)
		}

		pureIds[id] = true
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				if grid[i][j] > 0 {
					pureIds[id] = false
					pureIds[grid[i][j]] = false
				} else {
					grid[i][j] = id
				}
			}
		}
	}

	for id, pure := range pureIds {
		if pure {
			log.Printf("Winner: #%d", id)
		}
	}
}

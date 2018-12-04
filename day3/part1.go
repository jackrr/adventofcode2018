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

	collidedCells := 0

	var x, y, w, h int
	for _, box := range boxes {
		tokens := strings.Split(box, " ")
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

		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				grid[i][j] = grid[i][j] + 1
				if grid[i][j] == 2 {
					collidedCells++
				}
			}
		}
	}

	log.Printf("Found %d collisions", collidedCells)
}

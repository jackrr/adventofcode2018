package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	inputString := string(input)
	frequency := 0
	var nextNumString strings.Builder
	pastValues := make(map[int]bool)
	for {
		for _, c := range inputString {
			if c == '\n' {
				prev := frequency
				delta, err := strconv.Atoi(nextNumString.String())

				if err != nil {
					panic(err)
				}

				frequency = prev + delta
				if pastValues[frequency] {
					log.Printf("Already seen %d.", frequency)
					os.Exit(0)
				}
				pastValues[frequency] = true

				log.Printf("Current: %d, delta: %d, result: %d", prev, delta, frequency)
				nextNumString.Reset()
				continue
			}

			nextNumString.WriteRune(c)
		}
	}

	log.Printf("Final frequency: %d", frequency)
}

package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	at    time.Time
	id    int
	sleep bool
	wake  bool
}

type Guard struct {
	id     int
	shifts [][]bool
}

func (e *Event) Before(other *Event) bool {
	return e.at.Before(other.at)
}

func sort(events []Event) []Event {
	if len(events) < 2 {
		return events
	}

	left := make([]Event, 0)
	right := make([]Event, 0)
	for idx, event := range events {
		if idx < len(events)/2 {
			left = append(left, event)
		} else {
			right = append(right, event)
		}
	}

	left = sort(left)
	right = sort(right)

	var leftIdx, rightIdx, sortedIdx int
	sorted := make([]Event, len(events))
	for leftIdx < len(left) || rightIdx < len(right) {
		if rightIdx >= len(right) {
			sorted[sortedIdx] = left[leftIdx]
			leftIdx++
			sortedIdx++
			continue
		}

		if leftIdx >= len(left) {
			sorted[sortedIdx] = right[rightIdx]
			rightIdx++
			sortedIdx++
			continue
		}

		if left[leftIdx].Before(&right[rightIdx]) {
			sorted[sortedIdx] = left[leftIdx]
			leftIdx++
		} else {
			sorted[sortedIdx] = right[rightIdx]
			rightIdx++
		}
		sortedIdx++
	}

	return sorted
}

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	eventStrings := strings.Split(string(input), "\n")
	events := make([]Event, len(eventStrings))
	for idx, str := range eventStrings {
		event := Event{}

		event.at, err = time.Parse("2006-01-02 15:04", str[1:17])
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(str, "Guard") {
			ids := str[26 : len(str)-13]
			event.id, err = strconv.Atoi(ids)
			if err != nil {
				log.Fatal(err)
			}
		}

		if strings.Contains(str, "wakes") {
			event.wake = true
		}

		if strings.Contains(str, "asleep") {
			event.sleep = true
		}

		events[idx] = event
	}

	sorted := sort(events)

	var gid, start, end int
	guards := make(map[int]Guard, 0)
	for _, e := range sorted {
		if e.id > 0 {
			gid = e.id
		}

		if e.sleep {
			start = e.at.Minute()
		}

		if e.wake {
			shift := make([]bool, 60)
			end = e.at.Minute()
			for i := start; i < end; i++ {
				shift[i] = true
			}

			guard := guards[gid]
			if guard.id == 0 {
				guard = Guard{id: gid, shifts: make([][]bool, 0)}
			}
			guard.shifts = append(guard.shifts, shift)
			guards[gid] = guard
		}
	}

	var max, chosenMinute int
	var chosen Guard
	for _, guard := range guards {
		count := 0
		minutes := make(map[int]int, 0)
		for _, shift := range guard.shifts {
			for minute, asleep := range shift {
				if asleep {
					minutes[minute]++
					count++
				}
			}
		}

		if count > max {
			max = count
			chosen = guard
			maxTimes := 0
			for minute, times := range minutes {
				if times > maxTimes {
					maxTimes = times
					chosenMinute = minute
				}
			}
		}
	}

	log.Printf("chose %v %d: %d", chosen.id, chosenMinute, chosen.id*chosenMinute)

}

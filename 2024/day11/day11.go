package day11

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day int = 11

type Lineup struct {
	// map of stone value => count
	stones map[int]int
}

func NewLineup() *Lineup {
	return &Lineup{
		stones: make(map[int]int, 0),
	}
}

func (lineup *Lineup) AddStone(stone int, increment int) {
	_, ok := lineup.stones[stone]
	if ok {
		lineup.stones[stone] += increment
	} else {
		lineup.stones[stone] = increment
	}
}

func (lineup *Lineup) RemoveStone(stone int, decrement int) {
	_, ok := lineup.stones[stone]
	if ok {
		if lineup.stones[stone]-decrement == 0 {
			delete(lineup.stones, stone)
		} else {
			lineup.stones[stone] -= decrement
		}
	}
}

func (lineup *Lineup) Blink() {
	// create a new lineup
	newLineup := NewLineup()

	for stone, count := range lineup.stones {
		newLineup.stones[stone] = count
	}

	// apply rules
	for stone, count := range lineup.stones {
		if stone == 0 {
			newLineup.AddStone(1, count)
			newLineup.RemoveStone(0, count)
			continue
		}
		valueStr := fmt.Sprintf("%d", stone)
		if len(valueStr)%2 == 0 {
			// even number of digits, split in half
			midPoint := len(valueStr) / 2
			left := input.GetInt(valueStr[0:midPoint])
			right := input.GetInt(valueStr[midPoint:])

			newLineup.AddStone(left, count)
			newLineup.AddStone(right, count)
			newLineup.RemoveStone(stone, count)

			continue
		}

		newLineup.AddStone(stone*2024, count)
		newLineup.RemoveStone(stone, count)
	}
	// reassign pointer
	*lineup = *newLineup
}

func (lineup *Lineup) Size() int {
	total := 0
	for _, count := range lineup.stones {
		total += count
	}
	return total
}

func (lineup *Lineup) String() string {
	values := make([]string, 0)

	for stone, count := range lineup.stones {
		values = append(values, fmt.Sprintf("Value: %d, count: %d", stone, count))
	}

	return strings.Join(values, "\n")
}

func Part1and2() {
	contents := input.GetFileContents(day)
	numbers := strings.Split(contents, " ")

	lineup := NewLineup()

	for _, number := range numbers {
		lineup.AddStone(input.GetInt(number), 1)
	}

	// part 1
	blinks := 25

	for counter := 0; counter < blinks; counter++ {
		// process line-up
		lineup.Blink()
	}

	fmt.Printf("Part 1 answer: %d\n", lineup.Size())

	// part 2
	blinks = 75

	for counter := 25; counter < blinks; counter++ {
		// process line-up
		lineup.Blink()
	}

	fmt.Printf("Part 2 answer: %d\n", lineup.Size())
}

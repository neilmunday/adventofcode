package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"flag"
	"fmt"
	"os"
)

func main() {
	var day int

	flag.IntVar(&day, "day", 0, "Specify which AOC day challenge to run")
	flag.Parse()

	if day == 0 {
		flag.Usage()
		os.Exit(1)
	}

	switch day {
	case 1:
		day1.Part1()
		day1.Part2()
	case 2:
		day2.Part1()
		day2.Part2()
	case 3:
		day3.Part1()
		day3.Part2()
	case 4:
		day4.Part1()
		day4.Part2()
	case 5:
		day5.Part1()
		day5.Part2()
	case 6:
		day6.Part1()
		day6.Part2()
	default:
		fmt.Printf("Invalid day: %d\n", day)
		os.Exit(1)
	}
}

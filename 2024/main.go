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

	if day == 1 {
		day1.Part1()
		day1.Part2()
	} else if day == 2 {
		day2.Part1()
		day2.Part2()
	} else if day == 3 {
		day3.Part1()
		day3.Part2()
	} else if day == 4 {
		day4.Part1()
		day4.Part2()
	} else if day == 5 {
		day5.Part1and2()
	} else if day == 6 {
		day6.Part1and2()
	} else {
		fmt.Printf("Invalid day: %d", day)
		os.Exit(1)
	}
}

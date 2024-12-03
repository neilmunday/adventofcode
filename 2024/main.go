package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"flag"
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
	}
}

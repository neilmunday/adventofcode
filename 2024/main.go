package main

import (
	"aoc/day1"
	"aoc/day10"
	"aoc/day11"
	"aoc/day12"
	"aoc/day13"
	"aoc/day14"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"aoc/day9"
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
	} else if day == 7 {
		day7.Part1and2()
	} else if day == 8 {
		day8.Part1()
		day8.Part2()
	} else if day == 9 {
		day9.Part1()
		day9.Part2()
	} else if day == 10 {
		day10.Part1()
		day10.Part2()
	} else if day == 11 {
		day11.Part1and2()
	} else if day == 12 {
		day12.Part1and2()
	} else if day == 13 {
		day13.Part1and2()
	} else if day == 14 {
		day14.Part1and2(false)
		day14.Part1and2(true)
	} else {
		fmt.Printf("Invalid day: %d\n", day)
		os.Exit(1)
	}
}

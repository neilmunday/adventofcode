package main

import (
	"aoc/day1"
	"aoc/day2"
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
	} else {
		fmt.Printf("Invalid day: %d\n", day)
		os.Exit(1)
	}
}

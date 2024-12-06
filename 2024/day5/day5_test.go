package day5

import "testing"

func TestPart1(t *testing.T) {
	part1Expected := 5452

	part1, _ := Part1and2()

	if part1 != part1Expected {
		t.Errorf("Part1: expected %d, got %d", part1Expected, part1)
	}
}

func TestPart2(t *testing.T) {
	part2Expected := 4598

	_, part2 := Part1and2()

	if part2 != part2Expected {
		t.Errorf("Part2: expected %d, got %d", part2Expected, part2)
	}
}

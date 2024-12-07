package day6

import "testing"

func TestPart1and2(t *testing.T) {
	part1Expected := 5318
	part2Expected := 1831

	part1, part2 := Part1and2()

	if part1 != part1Expected {
		t.Errorf("Part1: expected %d, got %d", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("Part2: expected %d, got %d", part2Expected, part2)
	}
}

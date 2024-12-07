package day1

import "testing"

func TestPart1(t *testing.T) {
	part1Expected := 2367773

	part1 := Part1()

	if part1 != part1Expected {
		t.Errorf("Part1: expected %d, got %d", part1Expected, part1)
	}
}

func TestPart2(t *testing.T) {
	part2Expected := 21271939

	part2 := Part2()

	if part2 != part2Expected {
		t.Errorf("Part2: expected %d, got %d", part2Expected, part2)
	}
}

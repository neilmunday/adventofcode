package day7

import "testing"

func TestPart1and2(t *testing.T) {
	const part1Expected int64 = 5030892084481
	const part2Expected int64 = 91377448644679

	part1, part2 := Part1and2()

	if part1 != part1Expected {
		t.Errorf("Part1: expected %d, got %d", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("Part2: expected %d, got %d", part2Expected, part2)
	}
}

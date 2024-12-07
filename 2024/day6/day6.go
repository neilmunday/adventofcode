package day6

import (
	"aoc/input"
	"fmt"
	"strings"
)

const day = 6

type coord struct {
	x int
	y int
}

type move struct {
	x         int
	y         int
	direction string
}

func replaceAtIndex(in string, char rune, i int) string {
	out := []rune(in)
	out[i] = char
	return string(out)
}

func setVisited(x int, y int, direction string, visitedMap map[coord]bool, visitedList *[]move) {
	c := coord{x: x, y: y}

	if !visitedMap[c] {
		*visitedList = append(*visitedList, move{x: x, y: y, direction: direction})
	}

	visitedMap[c] = true
}

func setPath(x int, y int, direction string, path map[move]bool) bool {
	m := move{x: x, y: y, direction: direction}

	if path[m] {
		return false
	}

	path[m] = true

	return true
}

func Part1and2() (int, int) {

	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	stop := false

	x := 0
	y := 0

	direction := ""

	// find starting position
	for index, line := range lines {
		pos := strings.Index(line, "^")
		if pos != -1 {
			x = pos
			y = index
			direction = string(line[pos])
			break
		}
		pos = strings.Index(line, ">")
		if pos != -1 {
			x = pos
			y = index
			direction = string(line[pos])
			break
		}
		pos = strings.Index(line, "<")
		if pos != -1 {
			x = pos
			y = index
			direction = string(line[pos])
			break
		}
		pos = strings.Index(line, "V")
		if pos != -1 {
			x = pos
			y = index
			direction = string(line[pos])
			break
		}
	}

	start_x := x
	start_y := y
	start_direction := direction

	height := len(lines)
	width := len(lines[0])

	visitedMap := make(map[coord]bool)
	visitedList := []move{}
	setVisited(x, y, direction, visitedMap, &visitedList)

	for {
		if stop {
			break
		}

		if direction == "^" {
			// go up until we hit an obstacle or the top edge and exit
			for {
				if y-1 >= 0 {
					if string(lines[y-1][x]) != "#" {
						y -= 1
						setVisited(x, y, direction, visitedMap, &visitedList)
					} else {
						direction = ">"
						break
					}
				} else {
					stop = true
					break
				}
			}
		} else if direction == "V" {
			// go down until we hit an obstacle or the top edge and exit
			for {
				if y+1 <= height-1 {
					if string(lines[y+1][x]) != "#" {
						y += 1
						setVisited(x, y, direction, visitedMap, &visitedList)
					} else {
						direction = "<"
						break
					}
				} else {
					stop = true
					break
				}
			}
		} else if direction == ">" {
			// go right until we hit an obstacle or the top edge and exit
			for {
				if x+1 <= width-1 {
					if string(lines[y][x+1]) != "#" {
						x += 1
						setVisited(x, y, direction, visitedMap, &visitedList)
					} else {
						direction = "V"
						break
					}
				} else {
					stop = true
					break
				}
			}
		} else if direction == "<" {
			// go left until we hit an obstacle or the top edge and exit
			for {
				if x-1 >= 0 {
					if string(lines[y][x-1]) != "#" {
						x -= 1
						setVisited(x, y, direction, visitedMap, &visitedList)
					} else {
						direction = "^"
						break
					}
				} else {
					stop = true
					break
				}
			}
		}
	}

	part1Answer := len(visitedMap)
	fmt.Printf("Part 1 answer: %d\n", part1Answer)

	// now place an object on each place visited (except the start)
	// and see if this would create a loop

	loops := 0

	// copy the input
	newLines := make([]string, len(lines))
	for index, line := range lines {
		newLines[index] = strings.Clone(line)
	}

	obstacle_x := -1
	obstacle_y := -1

	for visitNumber, visited := range visitedList[1:] {

		if visitNumber > 0 {
			// remove old obstacle
			newLines[obstacle_y] = replaceAtIndex(newLines[obstacle_y], '.', obstacle_x)
		}

		obstacle_x = visited.x
		obstacle_y = visited.y

		// place obstacle
		newLines[obstacle_y] = replaceAtIndex(newLines[obstacle_y], '#', obstacle_x)

		x = start_x
		y = start_y
		direction := start_direction

		path := make(map[move]bool)
		stop = false

		for {
			if stop {
				break
			}

			if direction == "^" {
				// go up until we hit an obstacle or the top edge and exit or loop
				for {
					if y-1 >= 0 {
						if string(newLines[y-1][x]) != "#" {
							y -= 1
							if !setPath(x, y, direction, path) {
								loops += 1
								stop = true
								break
							}
						} else {
							direction = ">"
							break
						}
					} else {
						stop = true
						break
					}
				}
			} else if direction == "V" {
				// go down until we hit an obstacle or the top edge and exit or loop
				for {
					if y+1 <= height-1 {
						if string(newLines[y+1][x]) != "#" {
							y += 1
							if !setPath(x, y, direction, path) {
								loops += 1
								stop = true
								break
							}
						} else {
							direction = "<"
							break
						}
					} else {
						stop = true
						break
					}
				}
			} else if direction == ">" {
				// go right until we hit an obstacle or the top edge and exit or loop
				for {
					if x+1 <= width-1 {
						if string(newLines[y][x+1]) != "#" {
							x += 1
							if !setPath(x, y, direction, path) {
								loops += 1
								stop = true
								break
							}
						} else {
							direction = "V"
							break
						}
					} else {
						stop = true
						break
					}
				}
			} else if direction == "<" {
				// go left until we hit an obstacle or the top edge and exit or loop
				for {
					if x-1 >= 0 {
						if string(newLines[y][x-1]) != "#" {
							x -= 1
							if !setPath(x, y, direction, path) {
								loops += 1
								stop = true
								break
							}
						} else {
							direction = "^"
							break
						}
					} else {
						stop = true
						break
					}
				}
			}
		}
	}

	fmt.Printf("Part 2 answer: %d\n", loops)

	return part1Answer, loops
}

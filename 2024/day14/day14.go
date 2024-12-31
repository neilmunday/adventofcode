package day14

import (
	"aoc/input"
	"fmt"
	"regexp"
	"strings"
)

const day int = 14

type Vector struct {
	x int
	y int
}

type Robot struct {
	pos        GridPoint
	startPos   GridPoint
	v          Vector
	gridWidth  int
	gridHeight int
}

func NewRobot(x int, y int, gridWidth int, gridHeight int, vector Vector) *Robot {
	return &Robot{
		pos: GridPoint{
			x: x,
			y: y,
		},
		startPos: GridPoint{
			x: x,
			y: y,
		},
		v:          vector,
		gridWidth:  gridWidth,
		gridHeight: gridHeight,
	}
}

func (robot *Robot) Move(seconds int) {
	// move robot using vector and no. of seconds
	robot.pos.x = (seconds * robot.v.x) + robot.startPos.x
	robot.pos.y = (seconds * robot.v.y) + robot.startPos.y

	if robot.pos.x < 0 || robot.pos.x > robot.gridWidth-1 {
		// gone out of bounds
		rem := robot.pos.x % robot.gridWidth

		if rem == 0 {
			robot.pos.x = 0
		} else if robot.pos.x < 0 {
			robot.pos.x = robot.gridWidth + rem
		} else {
			robot.pos.x = rem
		}
	}

	if robot.pos.y < 0 || robot.pos.y > robot.gridHeight-1 {
		// gone out of bounds
		rem := robot.pos.y % robot.gridHeight

		if rem == 0 {
			robot.pos.y = 0
		} else if robot.pos.y < 0 {
			robot.pos.y = robot.gridHeight + rem
		} else {
			robot.pos.y = rem
		}
	}
}

func (robot *Robot) Reset() {
	robot.pos.x = robot.startPos.x
	robot.pos.y = robot.startPos.y
}

func (robot *Robot) String() string {
	return fmt.Sprintf("{Robot (%d, %d)}", robot.pos.x, robot.pos.y)
}

type GridPoint struct {
	x int
	y int
}

type Grid struct {
	width    int
	height   int
	robots   []*Robot
	robotMap map[GridPoint][]*Robot
}

func NewGrid(width int, height int) *Grid {
	return &Grid{
		width:    width,
		height:   height,
		robots:   make([]*Robot, 0),
		robotMap: make(map[GridPoint][]*Robot),
	}
}

func (grid *Grid) AddRobot(robot *Robot) {
	grid.robots = append(grid.robots, robot)
	gp := GridPoint{
		x: robot.pos.x,
		y: robot.pos.y,
	}
	grid.robotMap[gp] = append(grid.robotMap[gp], robot)
}

func (grid *Grid) MoveRobots(seconds int) {
	// make a new map
	grid.robotMap = make(map[GridPoint][]*Robot)
	// move robots
	for _, robot := range grid.robots {
		robot.Move(seconds)
		gp := GridPoint{
			x: robot.pos.x,
			y: robot.pos.y,
		}
		grid.robotMap[gp] = append(grid.robotMap[gp], robot)
	}
}

func (grid *Grid) QuadrantCount() (int, int, int, int) {
	tl := 0
	tr := 0
	bl := 0
	br := 0

	midWidth := grid.width / 2
	midHeight := grid.height / 2

	for _, robot := range grid.robots {
		// tl
		if robot.pos.x >= 0 && robot.pos.x < midWidth && robot.pos.y >= 0 && robot.pos.y < midHeight {
			tl++
		}
		// tr
		if robot.pos.x > midWidth && robot.pos.x < grid.width && robot.pos.y >= 0 && robot.pos.y < midHeight {
			tr++
		}
		// bl
		if robot.pos.x >= 0 && robot.pos.x < midWidth && robot.pos.y > midHeight && robot.pos.y < robot.gridHeight {
			bl++
		}
		// br
		if robot.pos.x > midWidth && robot.pos.x < grid.width && robot.pos.y > midHeight && robot.pos.y < robot.gridHeight {
			br++
		}
	}

	return tl, tr, bl, br
}

func (grid *Grid) Reset() {
	grid.robotMap = make(map[GridPoint][]*Robot)
	for _, robot := range grid.robots {
		robot.Reset()
		grid.robotMap[robot.pos] = append(grid.robotMap[robot.pos], robot)
	}
}

func (grid *Grid) SafetyFactor() int {
	tl, tr, bl, br := grid.QuadrantCount()
	return tl * tr * bl * br
}

func (grid *Grid) String() string {
	s := ""
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			// is there a robot at this spot?
			gp := GridPoint{
				x: x,
				y: y,
			}

			robotsArray, ok := grid.robotMap[gp]

			if ok {
				s += fmt.Sprintf("%d", len(robotsArray))
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func Part1and2(part2 bool) {
	contents := input.GetFileContents(day)

	grid := NewGrid(101, 103)
	//grid := NewGrid(11, 7)

	lines := strings.Split(contents, "\n")

	lineRe := regexp.MustCompile(`^p=(-?[\d]+),(-?[\d]+) v=(-?[\d]+),(-?[\d]+)$`)

	for lineNumber, line := range lines {
		inputMatches := lineRe.FindStringSubmatch(line)

		if len(inputMatches) == 0 {
			panic(fmt.Sprintf("Invalid input on line %d\n", lineNumber+1))
		}

		robot := NewRobot(input.GetInt(inputMatches[1]), input.GetInt(inputMatches[2]), grid.width, grid.height, Vector{
			x: input.GetInt(inputMatches[3]),
			y: input.GetInt(inputMatches[4]),
		})

		grid.AddRobot(robot)
	}

	if !part2 {
		// part 1
		grid.MoveRobots(100)
		fmt.Printf("Part 1 answer: %d\n", grid.SafetyFactor())
	} else {
		// part2
		// let's look for a line of robots to find the tree..

		stop := false

		for i := 0; i < 10000; i++ {
			grid.MoveRobots(i)

			for _, robot := range grid.robots {
				found := 0
				for x := robot.pos.x; x < grid.width; x++ {
					gp := GridPoint{
						x: x,
						y: robot.pos.y,
					}

					_, ok := grid.robotMap[gp]

					if ok {
						found++
					} else {
						break
					}
				}

				if found >= 10 {
					fmt.Println(grid)
					stop = true
					break
				}
			}

			if stop {
				fmt.Printf("Part 2 answer: %d\n", i)
				break
			}
		}
	}

}

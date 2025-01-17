package day15

import (
	"aoc/input"
	"fmt"
	"sort"
	"strings"
)

const day int = 15

type GridValue int

const (
	Wall GridValue = iota
	Box
	Path
	Robot
	BoxLeft
	BoxRight
	Visited
)

type GridPoint struct {
	x     int
	y     int
	value GridValue
}

type Grid struct {
	width    int
	height   int
	robotPos *GridPoint
	points   [][]GridPoint
}

func NewGrid(width int, height int) *Grid {

	grid := Grid{
		width:    width,
		height:   height,
		robotPos: nil,
		points:   make([][]GridPoint, width),
	}

	for x := 0; x < width; x++ {
		grid.points[x] = make([]GridPoint, height)
		for y := 0; y < height; y++ {
			grid.points[x][y] = GridPoint{
				x:     x,
				y:     y,
				value: Wall,
			}
		}
	}

	return &grid
}

func (grid *Grid) GetBoxGPSSum() int {
	total := 0

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			if grid.points[x][y].value == Box {
				total += x + (y * 100)
			}
		}
	}

	return total
}

func (grid *Grid) MoveRobotHorizontal(xDiff int) {
	xDiffOps := xDiff * -1

	if grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value != Wall {
		if grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value == Box {
			start := grid.robotPos.x + xDiff
			x := start + xDiff
			canShift := false
			for {
				if grid.points[x][grid.robotPos.y].value == Path {
					// we can shift
					canShift = true
					break
				}
				if grid.points[x][grid.robotPos.y].value == Wall {
					// can't shift
					break
				}
				x += xDiff
			}

			if canShift {
				for {
					grid.points[x][grid.robotPos.y].value = grid.points[x+xDiffOps][grid.robotPos.y].value
					if x == start {
						break
					}
					x += xDiffOps
				}
				grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
				grid.robotPos = &grid.points[x][grid.robotPos.y]
			}
		} else if grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value == Path {
			grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
			grid.robotPos = &grid.points[grid.robotPos.x+xDiff][grid.robotPos.y]
			grid.points[grid.robotPos.x][grid.robotPos.y].value = Robot
		}
	}
}

func (grid *Grid) MoveRobotVertical(yDiff int) {
	yDiffOps := yDiff * -1

	if grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value != Wall {
		if grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value == Box {
			start := grid.robotPos.y + yDiff
			y := start + yDiff
			canShift := false
			for {
				if grid.points[grid.robotPos.x][y].value == Path {
					// we can shift
					canShift = true
					break
				}
				if grid.points[grid.robotPos.x][y].value == Wall {
					// can't shift
					break
				}
				y += yDiff
			}

			if canShift {
				for {
					grid.points[grid.robotPos.x][y].value = grid.points[grid.robotPos.x][y+yDiffOps].value
					if y == start {
						break
					}
					y += yDiffOps
				}
				grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
				grid.robotPos = &grid.points[grid.robotPos.x][y]
				grid.points[grid.robotPos.x][grid.robotPos.y].value = Robot
			}
		} else if grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value == Path {
			grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
			grid.robotPos = &grid.points[grid.robotPos.x][grid.robotPos.y+yDiff]
			grid.points[grid.robotPos.x][grid.robotPos.y].value = Robot
		}
	}
}

func (grid *Grid) MoveRobotDown() {
	grid.MoveRobotVertical(1)
}

func (grid *Grid) MoveRobotLeft() {
	grid.MoveRobotHorizontal(-1)
}

func (grid *Grid) MoveRobotRight() {
	grid.MoveRobotHorizontal(1)
}

func (grid *Grid) MoveRobotUp() {
	grid.MoveRobotVertical(-1)
}

func (grid *Grid) SetGridPoint(x int, y int, value GridValue) {
	grid.points[x][y].value = value

	if value == Robot {
		grid.robotPos = &grid.points[x][y]
	}
}

func (grid *Grid) String() string {
	s := ""

	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			switch value := grid.points[x][y].value; value {
			case Wall:
				s += "#"
			case Path:
				s += "."
			case Box:
				s += "O"
			case Robot:
				s += "@"
			}
		}
		s += "\n"
	}

	return s
}

type DoubleGrid struct {
	width    int
	height   int
	robotPos *GridPoint
	points   [][]GridPoint
}

func NewDoubleGrid(width int, height int) *DoubleGrid {

	grid := DoubleGrid{
		width:    width,
		height:   height,
		robotPos: nil,
		points:   make([][]GridPoint, width),
	}

	for x := 0; x < width; x++ {
		grid.points[x] = make([]GridPoint, height)
		for y := 0; y < height; y++ {
			if x == 0 || y == 0 || x == width-1 || y == height-1 {
				grid.points[x][y] = GridPoint{
					x:     x,
					y:     y,
					value: Wall,
				}
			} else {
				grid.points[x][y] = GridPoint{
					x:     x,
					y:     y,
					value: Path,
				}
			}
		}
	}

	return &grid
}

func (grid *DoubleGrid) GetBoxGPSSum() int {
	total := 0

	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			if grid.points[x][y].value == BoxLeft {
				total += x + (y * 100)
			}
		}
	}

	return total
}

func (grid *DoubleGrid) MoveRobotHorizontal(xDiff int) {
	xDiffOps := xDiff * -1

	if grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value != Wall {
		if grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value == BoxLeft || grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value == BoxRight {
			start := grid.robotPos.x + xDiff
			x := start + xDiff
			canShift := false
			for {
				if grid.points[x][grid.robotPos.y].value == Path {
					// we can shift
					canShift = true
					break
				}
				if grid.points[x][grid.robotPos.y].value == Wall {
					// can't shift
					break
				}
				x += xDiff
			}

			if canShift {
				for {
					grid.points[x][grid.robotPos.y].value = grid.points[x+xDiffOps][grid.robotPos.y].value
					if x == start {
						break
					}
					x += xDiffOps
				}
				grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
				grid.robotPos = &grid.points[x][grid.robotPos.y]
			}
		} else if grid.points[grid.robotPos.x+xDiff][grid.robotPos.y].value == Path {
			grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
			grid.robotPos = &grid.points[grid.robotPos.x+xDiff][grid.robotPos.y]
			grid.points[grid.robotPos.x][grid.robotPos.y].value = Robot
		}
	}
}

func (grid *DoubleGrid) MoveRobotVertical(yDiff int) {

	// wall blocking our path so return
	if grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value == Wall {
		return
	}
	// free space, let's move
	if grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value == Path {
		grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
		grid.robotPos = &grid.points[grid.robotPos.x][grid.robotPos.y+yDiff]
		grid.points[grid.robotPos.x][grid.robotPos.y].value = Robot
		return
	}

	if grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value != BoxLeft && grid.points[grid.robotPos.x][grid.robotPos.y+yDiff].value != BoxRight {
		panic(fmt.Sprintf("Unexpected value at: (%d, %d)", grid.robotPos.x, grid.robotPos.y+yDiff))
	}

	x := grid.robotPos.x
	y := grid.robotPos.y + yDiff

	pointsMap := make(map[GridPoint]bool)

	if grid.points[x][y].value == BoxLeft {
		if grid.Explore(x+1, y, yDiff, pointsMap) != Path {
			return
		}
	} else {
		if grid.Explore(x-1, y, yDiff, pointsMap) != Path {
			return
		}
	}

	rslt := grid.Explore(x, y, yDiff, pointsMap)

	if rslt == Wall {
		return
	}
	// can move boxes!
	// sort keys by y value

	gps := make([]GridPoint, 0, len(pointsMap))

	for gp := range pointsMap {
		gps = append(gps, gp)
	}

	if yDiff == -1 {
		sort.Slice(gps, func(i int, j int) bool {
			return gps[i].y < gps[j].y
		})
	} else {
		sort.Slice(gps, func(i int, j int) bool {
			return gps[i].y > gps[j].y
		})
	}

	// shift boxes
	for _, gp := range gps {
		grid.points[gp.x][gp.y+yDiff].value = grid.points[gp.x][gp.y].value
		grid.points[gp.x][gp.y].value = Path
	}
	// move robot
	grid.points[grid.robotPos.x][grid.robotPos.y].value = Path
	grid.robotPos = &grid.points[grid.robotPos.x][grid.robotPos.y+yDiff]
	grid.points[grid.robotPos.x][grid.robotPos.y].value = Robot
}

func (grid *DoubleGrid) Explore(x int, y int, yDiff int, pointsMap map[GridPoint]bool) GridValue {

	if grid.points[x][y].value == Path {
		return Path
	}

	pointsMap[grid.points[x][y]] = true

	if grid.points[x][y+yDiff].value == Wall {
		return Wall
	}

	if grid.points[x][y].value == BoxLeft && grid.points[x][y+yDiff].value == BoxLeft {
		return grid.Explore(x, y+yDiff, yDiff, pointsMap)
	}

	if grid.points[x][y].value == BoxRight && grid.points[x][y+yDiff].value == BoxRight {
		return grid.Explore(x, y+yDiff, yDiff, pointsMap)
	}

	if grid.points[x][y].value == BoxLeft && grid.points[x][y+yDiff].value == BoxRight || grid.points[x][y].value == BoxRight && grid.points[x][y+yDiff].value == BoxLeft {
		/*
			Straddle situation:
				[]
				 []
		*/
		l1 := grid.Explore(x, y+yDiff, yDiff, pointsMap)
		r1 := grid.Explore(x+1, y+yDiff, yDiff, pointsMap)
		l2 := grid.Explore(x, y+yDiff, yDiff, pointsMap)
		r2 := grid.Explore(x-1, y+yDiff, yDiff, pointsMap)

		if l1 == Wall || r1 == Wall || l2 == Wall || r2 == Wall {
			return Wall
		}

		if l1 == Path && r1 == Path && l2 == Path && r2 == Path {
			return Path
		}

		return Box
	}

	if grid.points[x][y].value == BoxRight && grid.points[x][y+yDiff].value == BoxLeft {
		return grid.Explore(x, y+yDiff, yDiff, pointsMap)
	}

	return Path
}

func (grid *DoubleGrid) MoveRobotDown() {
	grid.MoveRobotVertical(1)
}

func (grid *DoubleGrid) MoveRobotLeft() {
	grid.MoveRobotHorizontal(-1)
}

func (grid *DoubleGrid) MoveRobotRight() {
	grid.MoveRobotHorizontal(1)
}

func (grid *DoubleGrid) MoveRobotUp() {
	grid.MoveRobotVertical(-1)
}

func (grid *DoubleGrid) SanityCheck() bool {
	for x := 0; x < grid.width; x++ {
		for y := 0; y < grid.height; y++ {
			if grid.points[x][y].value == BoxLeft && grid.points[x+1][y].value != BoxRight {
				return false
			}
		}
	}
	return true
}

func (grid *DoubleGrid) SetGridPoint(x int, y int, value GridValue) {

	if value == Box {
		grid.points[x*2][y].value = BoxLeft
		grid.points[(x*2)+1][y].value = BoxRight
	} else if value == Robot {
		grid.points[x*2][y].value = Robot
		grid.points[(x*2)+1][y].value = Path
		grid.robotPos = &grid.points[x*2][y]
	} else {
		grid.points[x*2][y].value = value
		grid.points[(x*2)+1][y].value = value
	}
}

func (grid *DoubleGrid) String() string {
	s := ""

	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			switch value := grid.points[x][y].value; value {
			case Wall:
				s += "#"
			case Path:
				s += "."
			case BoxLeft:
				s += "["
			case BoxRight:
				s += "]"
			case Robot:
				s += "@"
			case Visited:
				s += "X"
			}
		}
		s += "\n"
	}

	return s
}

func Part1() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	// work out height first
	height := 0
	for i, line := range lines {
		if i == 0 {
			continue
		}
		height++
		if len(line) == 0 {
			break
		}
	}

	grid := NewGrid(len(lines[0]), height)

	directions := ""

	for y, line := range lines {
		if len(line) > 0 {
			if string(line[0]) == "#" {
				// parsing map
				for x, char := range line {
					switch char {
					case '#':
						grid.SetGridPoint(x, y, Wall)
					case '.':
						grid.SetGridPoint(x, y, Path)
					case 'O':
						grid.SetGridPoint(x, y, Box)
					case '@':
						grid.SetGridPoint(x, y, Robot)
					default:
						panic("Invalid char!")
					}
				}
			} else {
				directions += line
			}
		}
	}

	for i := range directions {
		switch char := directions[i]; char {
		case '<':
			grid.MoveRobotLeft()
		case '>':
			grid.MoveRobotRight()
		case '^':
			grid.MoveRobotUp()
		case 'v':
			grid.MoveRobotDown()
		}
	}

	fmt.Printf("Part 1 answer: %d\n", grid.GetBoxGPSSum())
}

func Part2() {
	contents := input.GetFileContents(day)
	lines := strings.Split(contents, "\n")

	// work out height first
	height := 0
	for i, line := range lines {
		if i == 0 {
			continue
		}
		height++
		if len(line) == 0 {
			break
		}
	}

	grid := NewDoubleGrid(len(lines[0])*2, height)

	directions := ""

	for y, line := range lines {
		if len(line) > 0 {
			if string(line[0]) == "#" {
				// parsing map
				for x, char := range line {
					switch char {
					case '#':
						grid.SetGridPoint(x, y, Wall)
					case '.':
						grid.SetGridPoint(x, y, Path)
					case 'O':
						grid.SetGridPoint(x, y, Box)
					case '@':
						grid.SetGridPoint(x, y, Robot)
					default:
						panic("Invalid char!")
					}
				}
			} else {
				directions += line
			}
		}
	}

	for i := range directions {
		switch char := directions[i]; char {
		case '<':
			grid.MoveRobotLeft()
		case '>':
			grid.MoveRobotRight()
		case '^':
			grid.MoveRobotUp()
		case 'v':
			grid.MoveRobotDown()
		}
	}

	//fmt.Println(grid)

	fmt.Printf("Part 2 answer: %d\n", grid.GetBoxGPSSum())
}

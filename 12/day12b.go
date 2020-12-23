package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Waypoint struct {
	north int
	east  int
}
func (w Waypoint) ToStr() string {
	desc := ""
	labelEast := ""
	labelNorth := ""
	if w.north < 0 {
		labelNorth = "South"
	} else {
		labelNorth = "North"
	}
	if w.east < 0 {
		labelEast = "West"
	} else {
		labelEast = "East"
	}

	desc = fmt.Sprintf("Waypoint is %d %s, %d %s", Abs(w.east), labelEast, Abs(w.north), labelNorth)

	return desc
}
func (w Waypoint) Rotate(instruction Instruction) Waypoint{
	if(instruction.action == "L" && instruction.value == 90) {
		w.east, w.north = -w.north, w.east
	}
	if(instruction.action == "R" && instruction.value == 90) {
		w.east, w.north = w.north, -w.east
	}
	if instruction.value == 180 {
		w.east, w.north = -w.east, -w.north
	}
	if(instruction.action == "L" && instruction.value == 270) {
		w.east, w.north = w.north, -w.east
	}
	if(instruction.action == "R" && instruction.value == 270) {
		w.east, w.north = -w.north, w.east
	}
	return w
}
type Boat struct {
	north     int
	east      int
}

func (b Boat) Move(w Waypoint,times int) Boat {
	b.east += w.east * times
	b.north += w.north * times
	return b
}

func (b Boat) ToStr() string {
	desc := ""
	labelEast := ""
	labelNorth := ""
	if b.north < 0 {
		labelNorth = "South"
	} else {
		labelNorth = "North"
	}
	if b.east < 0 {
		labelEast = "West"
	} else {
		labelEast = "East"
	}


	desc = fmt.Sprintf("Boat is %d %s, %d %s", Abs(b.east), labelEast, Abs(b.north), labelNorth)

	return desc
}

type Instruction struct {
	action string
	value  int
}

func (i Instruction) ToStr() string {
	desc := ""
	switch i.action {
	case "N":
		desc = fmt.Sprintf("Move waypoint North %d", i.value)
	case "S":
		desc = fmt.Sprintf("Move waypoint South %d", i.value)
	case "E":
		desc = fmt.Sprintf("Move waypoint East %d", i.value)
	case "W":
		desc = fmt.Sprintf("Move waypoint West %d", i.value)
	case "L":
		desc = fmt.Sprintf("Rotate waypoint counter-clockwise %d degrees", i.value)
	case "R":
		desc = fmt.Sprintf("Rotate waypoint clockwise %d degrees", i.value)
	case "F":
		desc = fmt.Sprintf("Move toward waypoint %d times", i.value)
	}
	return desc
}
func main() {
	var inputFilename string
	var waypoint Waypoint
	var boat Boat
	var instruction Instruction

	flag.StringVar(&inputFilename, "file", "nav.txt", "input file name")
	flag.Parse()

	waypoint = Waypoint{
		north: 1,
		east:  10,
	}
	boat = Boat{
		north:     0,
		east:      0,
	}
	fmt.Printf("\neast: %d north: %d\nWaypoint east: %d north: %d\n", boat.east, boat.north, waypoint.east, waypoint.north)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	for s.Scan() {

		line := s.Text()
		inst, val := parseInput(line)

		instruction = Instruction{
			action: inst,
			value:  val,
		}
		fmt.Printf("%s %s\t%s\n", boat.ToStr(),waypoint.ToStr(), instruction.ToStr())

		switch instruction.action {
		case "F":
			boat = boat.Move(waypoint,instruction.value)
		case "N":
			waypoint.north += val
		case "S":
			waypoint.north -= val
		case "W":
			waypoint.east -= val
		case "E":
			waypoint.east += val
		case "L":
			waypoint = waypoint.Rotate(instruction)
		case "R":
			waypoint = waypoint.Rotate(instruction)
		}
		fmt.Printf("\t%s\n\t%s\n", boat.ToStr(),waypoint.ToStr())

	}
	fmt.Printf("\nManhattan distance; %d + %d = %d\n", Abs(boat.east), Abs(boat.north), MD(boat.east,boat.north))

	err := s.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
func parseInput(line string) (string, int) {
	patStr := "^([NSEWLRF]{1})(.*)$"
	re := regexp.MustCompile(patStr)
	result := re.FindStringSubmatch(line)
	val, _ := strconv.ParseInt(result[2], 10, 32)
	return result[1], int(val)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MD(x int, y int) int {
	return Abs(x) + Abs(y)
}
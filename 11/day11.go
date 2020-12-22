package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

/*
	If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
	If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
	Otherwise, the seat's state does not change.
	Floor (.) never changes; seats don't move, and nobody sits on the floor.


	Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an
	occupied seat to become empty (rather than four or more from the previous rules).
	The other rules still apply: empty seats that see no occupied seats become occupied,
	seats matching no rule don't change, and floor never changes.
*/
type seat struct {
	row      int
	col      int
	value    string
	newvalue string
}

func (s seat) print() {
	fmt.Print(s.value)
}
func (s seat) isOccupied() bool {
	return s.value == "#"
}
func (s seat) isEmpty() bool {
	return s.value != "#"
}
func (s seat) isFloor() bool {
	return s.value == "."
}

func (s seat) adjacentEmpty() bool {
	return true
}

type seatrow []seat

func (sr seatrow) print() {
	for _, s := range sr {
		s.print()
	}
	fmt.Printf("\n")
}

type floor []seatrow

func (f floor) print() {
	for _, sr := range f {
		sr.print()
	}
	fmt.Printf("\n")

}

type room struct {
	rows int
	cols int
}

func (f floor) countOccupied() int {
	rows := len(f)
	cols := len(f[0])
	count := 0

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			//fmt.Printf("%+v\n", f[y][x])
			if f[y][x].isOccupied() {
				count++
			}
		}
	}
	return count
}
func (f floor) processSeat(s *seat) bool {
	if (*s).isFloor() {
		return false
	}
	if (*s).isEmpty() && f.isAdjacentEmpty(*s) {
		(*s).newvalue = "#"
		return true
	}
	if (*s).isOccupied() && f.numAdjacentOccupied(*s) >= 4 {
		(*s).newvalue = "L"
		return true
	}
	return false
}
func (f floor) processSeatAlt(s *seat) bool {
	// don't do anything with floor
	if (*s).isFloor() {
		return false
	}
	// if you look all ways and see no filled seats, sit in it
	if (*s).isEmpty() && f.canSeeEmpty(*s) {
		(*s).newvalue = "#"
		return true
	}
	// if you can see 5 filled seats, get up
	if (*s).isOccupied() && f.numSeenOccupied(*s) >= 5 {
		(*s).newvalue = "L"
		return true
	}
	return false
}
func (f floor) canSeeEmpty(s seat) bool {
	rv := true

	adjacent := f.getFirstTop(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstTopRight(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstRight(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstBottomRight(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstBottom(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstBottomLeft(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstLeft(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getFirstTopLeft(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	return rv
}
func (f floor) isAdjacentEmpty(s seat) bool {
	rv := true

	adjacent := f.getTop(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getTopRight(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getRight(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getBottomRight(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getBottom(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getBottomLeft(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getLeft(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	adjacent = f.getTopLeft(s)
	if adjacent != nil {
		rv = rv && (*adjacent).isEmpty()
	}
	if !rv {
		return rv
	}
	return rv
}
func (f floor) numAdjacentOccupied(s seat) int {
	rv := 0

	adjacent := f.getTop(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getTopRight(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getRight(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getBottomRight(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getBottom(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getBottomLeft(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getLeft(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getTopLeft(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	return rv
}
func (f floor) numSeenOccupied(s seat) int {
	rv := 0

	adjacent := f.getFirstTop(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstTopRight(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstRight(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstBottomRight(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstBottom(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstBottomLeft(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstLeft(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	adjacent = f.getFirstTopLeft(s)
	if adjacent != nil && (*adjacent).isOccupied() {
		rv++
	}
	return rv
}
func (f floor) getLeft(s seat) *seat {
	row := s.row
	col := s.col - 1
	if col >= 0 {
		return &f[row][col]
	}
	return nil
}
func (f floor) getRight(s seat) *seat {
	row := s.row
	col := s.col + 1
	if col < len(f[0]) {
		return &f[row][col]
	}
	return nil
}
func (f floor) getTop(s seat) *seat {
	row := s.row - 1
	col := s.col
	if row >= 0 {
		return &f[row][col]
	}
	return nil
}
func (f floor) getBottom(s seat) *seat {
	row := s.row + 1
	col := s.col
	if row < len(f) {
		return &f[row][col]
	}
	return nil
}
func (f floor) getTopLeft(s seat) *seat {
	row := s.row - 1
	col := s.col - 1
	if row >= 0 && col >= 0 {
		return &f[row][col]
	}
	return nil
}
func (f floor) getTopRight(s seat) *seat {
	row := s.row - 1
	col := s.col + 1
	if row >= 0 && col < len(f[0]) {
		return &f[row][col]
	}
	return nil
}
func (f floor) getBottomLeft(s seat) *seat {
	row := s.row + 1
	col := s.col - 1
	if row < len(f) && col >= 0 {
		return &f[row][col]
	}
	return nil
}
func (f floor) getBottomRight(s seat) *seat {
	row := s.row + 1
	col := s.col + 1
	if row < len(f) && col < len(f[0]) {
		return &f[row][col]
	}
	return nil
}

func (f floor) getFirstLeft(s seat) *seat {
	row := s.row
	col := s.col - 1
	for ; col >= 0; col-- {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
	}
	return nil
}
func (f floor) getFirstRight(s seat) *seat {
	row := s.row
	col := s.col + 1
	for ; col < len(f[0]); col++ {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
	}
	return nil
}
func (f floor) getFirstTop(s seat) *seat {
	row := s.row - 1
	col := s.col
	for ; row >= 0; row-- {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
	}
	return nil
}
func (f floor) getFirstBottom(s seat) *seat {
	row := s.row + 1
	col := s.col
	for ; row < len(f); row++ {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
	}
	return nil
}
func (f floor) getFirstTopLeft(s seat) *seat {
	row := s.row - 1
	col := s.col - 1
	for ; row >= 0 && col >= 0; {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
		col--
		row--
	}
	return nil
}
func (f floor) getFirstTopRight(s seat) *seat {
	row := s.row - 1
	col := s.col + 1
	for ; row >= 0 && col < len(f[0]); {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
		row--
		col++
	}
	return nil
}
func (f floor) getFirstBottomLeft(s seat) *seat {
	row := s.row + 1
	col := s.col - 1
	for ; row < len(f) && col >= 0; {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
		row++
		col--
	}
	return nil
}
func (f floor) getFirstBottomRight(s seat) *seat {
	row := s.row + 1
	col := s.col + 1
	for ; row < len(f) && col < len(f[0]); {
		if !f[row][col].isFloor() {
			return &f[row][col]
		}
		row++
		col++
	}
	return nil
}
func (f floor) saveChanges() {
	rows := len(f)
	cols := len(f[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			//fmt.Printf("%+v\n", f[y][x])
			if f[y][x].newvalue != "" {
				f[y][x].value = f[y][x].newvalue
				f[y][x].newvalue = ""
			}
		}
	}
}

func main() {
	var inputFilename string
	var seating floor

	flag.StringVar(&inputFilename, "file", "map.txt", "input file name")
	flag.Parse()
	seating = make(floor, 0)
	fp := openFile(inputFilename)
	s := bufio.NewScanner(fp)
	rowNum := 0
	for s.Scan() {

		line := s.Text()
		row := make(seatrow, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = seat{
				row:   rowNum,
				col:   i,
				value: string(line[i]),
			}
		}
		rowNum++
		seating = append(seating, row)
	}
	thisRoom := room{
		rows: len(seating),
		cols: len(seating[0]),
	}

	fmt.Println(thisRoom)
	seating.print()

	for {
		hasChanges := false

		for y := 0; y < thisRoom.rows; y++ {
			for x := 0; x < thisRoom.cols; x++ {
				rv := seating.processSeatAlt(&seating[y][x])
				hasChanges = hasChanges || rv
			}
		}
		//fmt.Printf("changes: %v", hasChanges)
		if hasChanges {
			seating.saveChanges()
			seating.print()
		} else {
			break

		}
	}
	fmt.Printf("Occupied seats: %d\n", seating.countOccupied())
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

//func adjacent(row int, col int) {
//	lt := [row - 1]
//}
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

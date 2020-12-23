package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	total := Abs(-5)
	if total != 5 {
		t.Errorf("Abs was incorrect, got: %d, want: %d.", total, 5)
	}
	total = Abs(5)
	if total != 5 {
		t.Errorf("Abs was incorrect, got: %d, want: %d.", total, 5)
	}
}

func TestWaypoint_Rotate(t *testing.T) {
	var w Waypoint
	w = Waypoint{east: 10, north: 5}
	nw := w.Rotate(Instruction{
		action: "L",
		value:  90,
	})
	if nw.east != -5 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nw.east, -5)
	}
	if nw.north != 10 {
		t.Errorf("North was incorrect, got: %d, want: %d.", nw.north, 10)
	}
	nw = w.Rotate(Instruction{
		action: "R",
		value:  90,
	})
	if nw.north != -10 {
		t.Errorf("North was incorrect, got: %d, want: %d.", nw.north, -10)
	}
	if nw.east != 5 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nw.east, 5)
	}
	nw = w.Rotate(Instruction{
		action: "R",
		value:  180,
	})
	if nw.north != -5 {
		t.Errorf("North was incorrect, got: %d, want: %d.", nw.north, -5)
	}
	if nw.east != -10 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nw.east, -10)
	}
	nw = w.Rotate(Instruction{
		action: "L",
		value:  180,
	})
	if nw.north != -5 {
		t.Errorf("North was incorrect, got: %d, want: %d.", nw.north, -5)
	}
	if nw.east != -10 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nw.east, -10)
	}
	nw = w.Rotate(Instruction{
		action: "R",
		value:  270,
	})
	if nw.north != 10 {
		t.Errorf("North was incorrect, got: %d, want: %d.", nw.north, 10)
	}
	if nw.east != -5 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nw.east, -5)
	}
	nw = w.Rotate(Instruction{
		action: "L",
		value:  270,
	})
	if nw.north != -10 {
		t.Errorf("North was incorrect, got: %d, want: %d.", nw.north, -10)
	}
	if nw.east != 5 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nw.east, 5)
	}
}
func TestBoat_Move(t *testing.T) {
	var w Waypoint
	var b Boat

	w = Waypoint{east: 10, north: 5}
	b = Boat{
		east:  0,
		north: 0,
	}
	nb := b.Move(w, 1)
	if nb.east != 10 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.east, 10)
	}
	if nb.north != 5 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.north, 5)
	}
	nb = b.Move(w, 2)
	if nb.east != 20 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.east, 20)
	}
	if nb.north != 10 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.north, 10)
	}

	w = Waypoint{east: 10, north: 5}
	b = Boat{
		east:  -10,
		north: -5,
	}
	nb = b.Move(w, 1)
	if nb.east != 0 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.east, 0)
	}
	if nb.north != 0 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.north, 0)
	}

	w = Waypoint{east: -10, north: 5}
	b = Boat{
		east:  -10,
		north: -5,
	}
	nb = b.Move(w, 1)
	if nb.east != -20 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.east, -20)
	}
	if nb.north != 0 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.north, 0)
	}


	w = Waypoint{east: -10, north: -15}
	b = Boat{
		east:  5,
		north: -5,
	}
	nb = b.Move(w, 3)
	if nb.east != -25 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.east, -25)
	}
	if nb.north != -50 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.north, -50)
	}


	w = Waypoint{east: 10, north: 15}
	b = Boat{
		east:  5,
		north: -5,
	}
	nb = b.Move(w, 3)
	if nb.east != 35 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.east, 35)
	}
	if nb.north != 40 {
		t.Errorf("East was incorrect, got: %d, want: %d.", nb.north, 40)
	}
}
func TestMD(t *testing.T) {
	r:=MD(-1,-1)

	if r != 2 {
		t.Errorf("East was incorrect, got: %d, want: %d.", r, 2)
	}
}
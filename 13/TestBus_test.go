package main

import (
	"testing"
)

func TestArrivalAfter(t *testing.T) {
	b := Bus {
		id: 4,
	}
	time := b.arrivalAfter(10)
	if time != 12 {
		t.Errorf("arrivalAfter was incorrect, got: %d, want: %d.", time, 12)

	}
	wait := b.waitTime(10)
	if wait != 2 {
		t.Errorf("waitTIme was incorrect, got: %d, want: %d.", wait, 2)

	}
	b.id = 7
	time = b.arrivalAfter(939)
	if time != 945 {
		t.Errorf("arrivalAfter was incorrect, got: %d, want: %d.", time, 945)

	}
	wait = b.waitTime(939)
	if wait != 6 {
		t.Errorf("arrivalAfter was incorrect, got: %d, want: %d.", wait, 6)

	}
}

func TestInverse(t *testing.T) {
	i:= inverse(17,29)
	if i != 12 {
		t.Errorf("Inverse was incorrect, got: %d, want: %d.", i, 12)
	}
}

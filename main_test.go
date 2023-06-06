package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		finishedOrderedList = []string{}
		dine()

		if len(finishedOrderedList) != len(philosophers) {
			t.Errorf("Different size of finished order list")
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", 0 * time.Second},
		{"quarter second delay", 250 * time.Millisecond},
		{"half second delay", 500 * time.Millisecond},
	}

	for _, e := range theTests {

		finishedOrderedList = []string{}

		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay
		dine()

		if len(finishedOrderedList) != len(philosophers) {
			t.Errorf("%s: Different size of finished order list; expected %d but got %d", e.name, len(philosophers), len(finishedOrderedList))
		}
	}

}

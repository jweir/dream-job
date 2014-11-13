package main

import (
	"testing"
	"time"
)

func TestScheduleExpasion(t *testing.T) {
	s := &Schedule{
		When: "0 */5 * * * * *",
	}

	r := s.expand(time.Now(), time.Minute*10)

	if len(r) != 2 {
		t.Fail()
		t.Logf("%d is not the expected size %v", len(r), r)
	}
}

func TestScheduleExpasionForPast(t *testing.T) {
	s := &Schedule{
		When: "0 */5 * * * * 1973",
	}

	r := s.expand(time.Now(), time.Minute*10)

	if len(r) != 0 {
		t.Fail()
	}
}

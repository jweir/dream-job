package main

import (
	"log"
	"time"

	"github.com/gorhill/cronexpr"
)

type Schedule struct {
	When string
	Job  *Job
	cron *cronexpr.Expression
}

type Job struct {
	Name string
}

func (s *Schedule) expand(t time.Time, w time.Duration) (times []time.Time) {
	until := t.Add(w).Unix()
	i := t.Unix()

	for i < until {
		next := s.cron.Next(time.Unix(i, 0))
		u := next.Unix()

		// If this is only scheduled for the past, break out
		if u < i {
			return
		}

		// This time is out of range
		if u > until {
			return
		}

		times = append(times, next)
		i = next.Unix()
	}

	return
}

func (s *Schedule) Start(suc chan (*Schedule)) {
	for {
		d := s.cron.Next(time.Now()).Sub(time.Now())
		time.Sleep(d)
		suc <- s
	}
}

func NewSchedule(when string, j *Job) *Schedule {
	return &Schedule{
		When: when,
		Job:  j,
		cron: cronexpr.MustParse(when),
	}
}

func main() {
	suc := make(chan (*Schedule))
	S := NewSchedule("*/5 * * * * * *", &Job{Name: "A"})
	go S.Start(suc)

	X := NewSchedule("*/30 * * * * * *", &Job{Name: "B"})
	go X.Start(suc)

	for {
		select {
		case event := <-suc:
			log.Printf("got message from %s", event.Job.Name)
		}
	}
}

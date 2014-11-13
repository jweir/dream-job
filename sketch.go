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

func (s *Schedule) Start() {
	for {
		d := s.cron.Next(time.Now()).Sub(time.Now())
		time.Sleep(d)
		log.Printf("tick %s\n", s.Job.Name)
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
	S := NewSchedule("*/5 * * * * * *", &Job{Name: "A"})
	go S.Start()

	X := NewSchedule("*/30 * * * * * *", &Job{Name: "B"})
	go X.Start()

	for {
		time.Sleep(time.Minute)
	}
}

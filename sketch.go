package main

import (
	"github.com/gorhill/cronexpr"
	"log"
	"time"
)

type CronString string
type JobArgs map[string]string

type Schedule struct {
	When      *cronexpr.Expression
	Job       *Job
	DoNotSkip bool
}

type Job struct {
	Name    string
	Perform func(JobArgs) error
}

func (s *Schedule) expand(t time.Time, w time.Duration) (times []time.Time) {
	until := t.Add(w).Unix()
	i := t.Unix()

	for i <= until {
		next := s.When.Next(time.Unix(i, 0))

		// If this is only scheduled for the past, break out
		if next.Unix() < i {
			i = until + 1
		} else {
			times = append(times, next)
			i = next.Unix()
		}
	}

	return
}

func main() {
	MyJ := &Job{
		Name: "Test",
		Perform: func(a JobArgs) error {
			log.Println("perform!")
			return nil
		},
	}

	S := &Schedule{
		When:      cronexpr.MustParse("10,20 5,10,30 */1 * * * *"),
		Job:       MyJ,
		DoNotSkip: false,
	}

	a := JobArgs{
		"a": "b",
	}

	S.Job.Perform(a)

	r := S.expand(time.Now(), time.Hour)

	for _, t := range r {
		log.Printf("%s\n", t)
	}
}

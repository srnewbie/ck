package cron

import (
	"math/rand"
	"time"
)

type (
	Cron interface {
		Start()
	}
	cronImpl struct {
		runners map[string]*runner
	}
	Config struct {
		Name     string
		Interval int
		Cb       func()
	}
	runner struct {
		name     string
		interval int
		previous time.Time
		timer    *time.Timer
		callback func()
	}
)

func New(c []*Config) Cron {
	cron := &cronImpl{runners: make(map[string]*runner)}
	for _, cfg := range c {
		cron.runners[cfg.Name] = &runner{
			name:     cfg.Name,
			interval: cfg.Interval,
			timer:    time.NewTimer(0),
			callback: cfg.Cb,
		}
	}
	return cron
}

func (c *cronImpl) Start() {
	for _, r := range c.runners {
		go r.run()
	}
}

func (r *runner) next() {
	if r.interval == -1 {
		rand.Seed(time.Now().UnixNano())
		r.interval = rand.Intn(10)
	}
	next := r.previous.Add(time.Duration(r.interval) * time.Second)
	r.timer.Reset(next.Sub(time.Now()))
}

func (r *runner) run() {
	for {
		<-r.timer.C
		r.previous = time.Now()
		r.callback()
		r.next()
	}
}

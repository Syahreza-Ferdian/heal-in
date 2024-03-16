package scheduler

import (
	"time"

	"github.com/robfig/cron/v3"
)

type SchedulerInterface interface {
	AddFunction(expression string, job func())
	Start()
	Stop()
}

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler(location *time.Location) SchedulerInterface {
	return &Scheduler{
		cron: cron.New(cron.WithLocation(location)),
	}
}

func (s *Scheduler) AddFunction(expression string, job func()) {
	s.cron.AddFunc(expression, job)
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	defer s.cron.Stop()
}
